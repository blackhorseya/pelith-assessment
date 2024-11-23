package etherscan

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TransactionRepoImpl is the implementation of TransactionRepo.
type TransactionRepoImpl struct {
	etherscanAPI *etherscan.Client
	ethclientAPI *ethclient.Client

	mu sync.Mutex

	abis map[string]abi.ABI
}

// NewTransactionRepoImpl is used to create a new TransactionRepoImpl.
func NewTransactionRepoImpl(app *configx.Application) (*TransactionRepoImpl, error) {
	etherscanAPI := etherscan.New(etherscan.Mainnet, app.Etherscan.APIKey)

	ethclientAPI, err := ethclient.Dial("https://mainnet.infura.io/v3/" + app.Infura.ProjectID)
	if err != nil {
		return nil, err
	}

	return &TransactionRepoImpl{
		etherscanAPI: etherscanAPI,
		ethclientAPI: ethclientAPI,
		mu:           sync.Mutex{},
		abis:         make(map[string]abi.ABI),
	}, nil
}

// NewTransactionGetter is used to create a new TransactionGetter.
func NewTransactionGetter(impl *TransactionRepoImpl) query.TransactionGetter {
	return impl
}

func (i *TransactionRepoImpl) ListByAddress(
	c context.Context,
	address string,
	cond query.ListTransactionCondition,
) (item biz.TransactionList, total int, err error) {
	ctx := contextx.WithContext(c)

	// 获取区块范围
	startBlock, err := i.etherscanAPI.BlockNumber(cond.StartTime.Unix(), "after")
	if err != nil {
		ctx.Error("failed to fetch start block", zap.Error(err), zap.Time("start_time", cond.StartTime))
		return nil, 0, err
	}

	if cond.EndTime.After(time.Now()) {
		cond.EndTime = time.Now()
	}
	endBlock, err := i.etherscanAPI.BlockNumber(cond.EndTime.Unix(), "before")
	if err != nil {
		ctx.Error("failed to fetch end block", zap.Error(err), zap.Time("end_time", cond.EndTime))
		return nil, 0, err
	}

	// 获取交易列表
	txs, err := i.etherscanAPI.NormalTxByAddress(address, &startBlock, &endBlock, 1, 100, true)
	if err != nil {
		ctx.Error("failed to fetch transactions", zap.Error(err))
		return nil, 0, err
	}

	var res biz.TransactionList
	var parsedABI abi.ABI
	var swapEventHash common.Hash

	// 获取目标合约的 ABI
	if cond.PoolAddress != "" {
		parsedABI, err = i.getABI(cond.PoolAddress)
		if err != nil {
			ctx.Error("failed to fetch contract ABI", zap.Error(err), zap.String("contract_address", cond.PoolAddress))
			return nil, 0, err
		}
		swapEventHash = parsedABI.Events["Swap"].ID
	}

	// 遍历交易并解析日志
	for _, tx := range txs {
		txType := model.TransactionType_TRANSACTION_TYPE_UNSPECIFIED
		var swapDetails []*model.SwapDetail

		if cond.PoolAddress != "" {
			// 获取交易的 Receipt
			receipt, err2 := i.ethclientAPI.TransactionReceipt(context.Background(), common.HexToHash(tx.Hash))
			if err2 != nil {
				ctx.Error("failed to fetch transaction receipt", zap.Error(err2), zap.String("tx_hash", tx.Hash))
				return nil, 0, err2
			}

			// 遍历日志，查找 Swap 事件
			for _, logEntry := range receipt.Logs {
				if !strings.EqualFold(logEntry.Address.Hex(), cond.PoolAddress) {
					continue // 非目标合约的日志
				}
				if len(logEntry.Topics) == 0 || logEntry.Topics[0] != swapEventHash {
					continue // 非 Swap 事件
				}

				// 提取日志的参数
				eventData := make(map[string]interface{})
				err = parsedABI.UnpackIntoMap(eventData, "Swap", logEntry.Data)
				if err != nil {
					ctx.Error("failed to unpack event data", zap.Error(err), zap.String("log_address", logEntry.Address.Hex()))
					continue
				}

				// 提取事件参数并构造 SwapDetails
				fromAddress := common.HexToAddress(logEntry.Topics[1].Hex())
				toAddress := common.HexToAddress(logEntry.Topics[2].Hex())

				fromTokenAmount, ok := eventData["amount0In"].(*big.Int)
				if !ok {
					ctx.Warn("missing or invalid amount0In", zap.Any("event_data", eventData))
					continue
				}
				toTokenAmount, ok := eventData["amount1Out"].(*big.Int)
				if !ok {
					ctx.Warn("missing or invalid amount1Out", zap.Any("event_data", eventData))
					continue
				}

				swapDetails = append(swapDetails, &model.SwapDetail{
					FromTokenAddress: fromAddress.Hex(),
					ToTokenAddress:   toAddress.Hex(),
					FromTokenAmount:  fromTokenAmount.Int64(),
					ToTokenAmount:    toTokenAmount.Int64(),
					PoolAddress:      cond.PoolAddress,
				})
			}

			// 如果找到 Swap 日志，标记交易类型
			if len(swapDetails) > 0 {
				txType = model.TransactionType_TRANSACTION_TYPE_SWAP
			}
		}

		// 构造 Transaction 实例
		res = append(res, &biz.Transaction{
			Transaction: model.Transaction{
				TxHash:      tx.Hash,
				FromAddress: tx.From,
				ToAddress:   tx.To,
				Amount:      tx.Value.Int().Int64(),
				Timestamp:   timestamppb.New(tx.TimeStamp.Time()),
				TaskId:      nil,
				CampaignId:  nil,
				Status:      model.TransactionStatus_TRANSACTION_STATUS_COMPLETED,
				Type:        txType,
				SwapDetails: swapDetails,
			},
		})
	}

	return res, len(res), nil
}

func (i *TransactionRepoImpl) getABI(contractAddress string) (abi.ABI, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Check cache first
	if a, ok := i.abis[contractAddress]; ok {
		return a, nil
	}

	// Fetch ABI JSON from etherscan API
	abiJSON, err := i.etherscanAPI.ContractABI(contractAddress)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to fetch ABI from Etherscan API: %w", err)
	}

	// Parse the ABI JSON
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse ABI JSON: %w", err)
	}

	// Cache the parsed ABI
	i.abis[contractAddress] = parsedABI

	return parsedABI, nil
}
