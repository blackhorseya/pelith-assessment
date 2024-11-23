package etherscan

import (
	"context"
	"fmt"
	"log"
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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

			// 解析 Swap 日志
			swapDetail, err2 := i.decodeSwapLogs(receipt.Logs, swapEventHash)
			if err2 != nil {
				ctx.Warn("failed to decode swap logs", zap.Error(err2), zap.String("tx_hash", tx.Hash))
				swapDetail = nil
			}

			if swapDetail != nil {
				txType = model.TransactionType_TRANSACTION_TYPE_SWAP
				swapDetails = append(swapDetails, swapDetail)
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

func (i *TransactionRepoImpl) decodeSwapLogs(logs []*types.Log, swapEventHash common.Hash) (*model.SwapDetail, error) {
	var firstLog, lastLog *types.Log
	var _, _ string
	var fromDecimals, toDecimals int
	var fromAmountFloat, toAmountFloat *big.Float

	// Iterate over logs to find the first and last valid Swap logs
	for _, logEntry := range logs {
		// Skip logs that don't match the criteria
		if len(logEntry.Topics) < 3 || logEntry.Topics[0] != swapEventHash {
			continue
		}

		// Ensure data length is sufficient
		if len(logEntry.Data) < 64 {
			return nil, fmt.Errorf("log data length is insufficient: %s", logEntry.Data)
		}

		// Set the first valid log if not already set
		if firstLog == nil {
			firstLog = logEntry
		}
		// Update the last valid log
		lastLog = logEntry
	}

	// Ensure we found at least one valid log
	if firstLog == nil || lastLog == nil {
		return nil, fmt.Errorf("no valid Swap log found")
	}

	// Parse the first log for "from" token details
	fromTokenAddress := firstLog.Address
	fromAmount := new(big.Int).SetBytes(firstLog.Data[:32]) // First 32 bytes represent the amount
	_, fromDecimals, err := i.getTokenDetails(fromTokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get From Token details (address: %s): %w", fromTokenAddress.Hex(), err)
	}
	fromAmountFloat = normalizeAmount(fromAmount, fromDecimals)

	// Parse the last log for "to" token details
	toTokenAddress := lastLog.Address
	toAmount := new(big.Int).SetBytes(lastLog.Data[:32]) // First 32 bytes represent the amount
	_, toDecimals, err = i.getTokenDetails(toTokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get To Token details (address: %s): %w", toTokenAddress.Hex(), err)
	}
	toAmountFloat = normalizeAmount(toAmount, toDecimals)

	return &model.SwapDetail{
		FromTokenAddress: fromTokenAddress.Hex(),
		ToTokenAddress:   toTokenAddress.Hex(),
		FromTokenAmount:  fromAmountFloat.String(),
		ToTokenAmount:    toAmountFloat.String(),
		PoolAddress:      "",
	}, nil
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

// getTokenDetails 获取 ERC20 Token 的 symbol 和 decimals
func (i *TransactionRepoImpl) getTokenDetails(tokenAddress common.Address) (string, int, error) {
	// ERC20 ABI，仅包含 symbol 和 decimals 方法
	const erc20ABI = `[{
		"constant": true,
		"inputs": [],
		"name": "symbol",
		"outputs": [{"name": "", "type": "string"}],
		"type": "function"
	}, {
		"constant": true,
		"inputs": [],
		"name": "decimals",
		"outputs": [{"name": "", "type": "uint8"}],
		"type": "function"
	}]`

	// 解析 ERC20 ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return "", 0, fmt.Errorf("解析 ERC20 ABI 失败: %v", err)
	}

	// 绑定 ERC20 合约
	contract := bind.NewBoundContract(tokenAddress, parsedABI, i.ethclientAPI, i.ethclientAPI, i.ethclientAPI)

	// 调用 symbol 方法
	var symbol string
	output := []interface{}{&symbol} // 包装输出为 []interface{}
	err = contract.Call(nil, &output, "symbol")
	if err != nil {
		log.Printf("无法获取 symbol (地址: %s): %v", tokenAddress.Hex(), err)
		symbol = "UNKNOWN" // 提供默认值
	}

	// 调用 decimals 方法
	var decimals uint8
	output = []interface{}{&decimals} // 重用切片包装 decimals
	err = contract.Call(nil, &output, "decimals")
	if err != nil {
		return symbol, 0, fmt.Errorf("无法获取 decimals (地址: %s): %v", tokenAddress.Hex(), err)
	}

	return symbol, int(decimals), nil
}

// 将 Token 金额根据 decimals 归一化为浮点数
func normalizeAmount(amount *big.Int, decimals int) *big.Float {
	// 计算 10^decimals
	decimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	// 将整数金额转换为浮点数，并归一化
	amountFloat := new(big.Float).SetInt(amount)           // 转换为浮点数
	decimalsFloat := new(big.Float).SetInt(decimalsFactor) // 转换为浮点数
	return new(big.Float).Quo(amountFloat, decimalsFloat)  // 执行归一化
}
