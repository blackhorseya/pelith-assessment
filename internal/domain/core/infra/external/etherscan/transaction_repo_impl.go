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

func (i *TransactionRepoImpl) GetByHash(c context.Context, hash string) (item *biz.Transaction, err error) {
	return i.getByHash(c, hash)
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
		var receipt *types.Receipt
		// 获取交易的 Receipt
		receipt, err = i.ethclientAPI.TransactionReceipt(context.Background(), common.HexToHash(tx.Hash))
		if err != nil {
			ctx.Error("failed to fetch transaction receipt", zap.Error(err), zap.String("tx_hash", tx.Hash))
			return nil, 0, err
		}

		// 构造 Transaction 实例
		got := biz.NewTransaction(
			tx.Hash,
			tx.From,
			tx.To,
			receipt.BlockNumber.Int64(),
			tx.TimeStamp.Time(),
		).WithReceipt(receipt)

		// 如果有目标合约，则解析日志
		if cond.PoolAddress != "" {
			swapDetail, err2 := got.GetSwapForPool(common.HexToHash(cond.PoolAddress), swapEventHash)
			if err2 != nil || swapDetail == nil {
				ctx.Debug(
					"the tx is not a swap tx",
					zap.String("tx_hash", got.GetTransaction().TxHash),
					zap.String("pool_address", cond.PoolAddress),
					zap.Error(err2),
				)
				continue
			}
		}

		res = append(res, got)
	}

	return res, len(res), nil
}

func (i *TransactionRepoImpl) GetLogsByAddress(
	c context.Context,
	contractAddress string,
	cond query.GetLogsCondition,
) (item []string, total int, err error) {
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

	parsedABI, err := i.getABI(contractAddress)
	if err != nil {
		ctx.Error(
			"failed to fetch contract ABI",
			zap.Error(err),
			zap.String("contract_address", contractAddress),
		)
		return nil, 0, err
	}
	swapEventHash := parsedABI.Events["Swap"].ID

	logs, err := i.etherscanAPI.GetLogs(startBlock, endBlock, contractAddress, swapEventHash.Hex())
	if err != nil {
		ctx.Error("failed to fetch logs", zap.Error(err))
		return nil, 0, err
	}

	for _, logEntry := range logs {
		item = append(item, logEntry.TransactionHash)
	}

	return item, len(item), nil
}

func (i *TransactionRepoImpl) GetSwapTxByPoolAddress(
	c context.Context,
	contractAddress string,
	cond query.ListTransactionCondition,
	txCh chan<- *biz.Transaction,
) error {
	ctx := contextx.WithContext(c)

	// 获取区块范围
	startBlock, err := i.etherscanAPI.BlockNumber(cond.StartTime.Unix(), "after")
	if err != nil {
		ctx.Error("failed to fetch start block", zap.Error(err), zap.Time("start_time", cond.StartTime))
		return err
	}

	if cond.EndTime.After(time.Now()) {
		cond.EndTime = time.Now()
	}
	endBlock, err := i.etherscanAPI.BlockNumber(cond.EndTime.Unix(), "before")
	if err != nil {
		ctx.Error("failed to fetch end block", zap.Error(err), zap.Time("end_time", cond.EndTime))
		return err
	}

	parsedABI, err := i.getABI(contractAddress)
	if err != nil {
		ctx.Error("failed to fetch contract ABI", zap.Error(err), zap.String("contract_address", contractAddress))
		return err
	}
	swapEventHash := parsedABI.Events["Swap"].ID

	logs, err := i.etherscanAPI.GetLogs(startBlock, endBlock, contractAddress, swapEventHash.Hex())
	if err != nil {
		ctx.Error("failed to fetch logs", zap.Error(err))
		return err
	}

	for _, logEntry := range logs {
		tx, err2 := i.getByHash(ctx, logEntry.TransactionHash)
		if err2 != nil || tx == nil {
			ctx.Error("failed to fetch transaction", zap.Error(err2), zap.String("tx_hash", logEntry.TransactionHash))
			return err2
		}

		swapDetail, err2 := tx.GetSwapForPool(common.HexToHash(contractAddress), swapEventHash)
		if err2 != nil || swapDetail == nil {
			ctx.Debug(
				"the tx is not a swap tx",
				zap.String("tx_hash", tx.GetTransaction().TxHash),
				zap.String("pool_address", contractAddress),
				zap.Error(err2),
			)
			continue
		}

		if txCh != nil {
			txCh <- tx
		}
	}

	return nil
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
		return "", 0, fmt.Errorf("解析 ERC20 ABI 失败: %w", err)
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
		return symbol, 0, fmt.Errorf("无法获取 decimals (地址: %s): %w", tokenAddress.Hex(), err)
	}

	return symbol, int(decimals), nil
}

func (i *TransactionRepoImpl) getFromByTx(ctx contextx.Contextx, tx *types.Transaction) (common.Address, error) {
	chainID, err := i.ethclientAPI.NetworkID(ctx)
	if err != nil {
		ctx.Error("failed to fetch network ID", zap.Error(err))
		return common.Address{}, err
	}

	signer := types.LatestSignerForChainID(chainID)
	from, err := types.Sender(signer, tx)
	if err != nil {
		ctx.Error("failed to fetch sender", zap.Error(err))
		return common.Address{}, err
	}

	return from, nil
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

func (i *TransactionRepoImpl) getByHash(c context.Context, hash string) (*biz.Transaction, error) {
	ctx := contextx.WithContext(c)

	// 获取交易
	tx, isPending, err := i.ethclientAPI.TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		ctx.Error("failed to fetch transaction", zap.Error(err), zap.String("tx_hash", hash))
		return nil, err
	}
	from, err := i.getFromByTx(ctx, tx)
	if err != nil {
		ctx.Error("failed to fetch sender", zap.Error(err), zap.String("tx_hash", hash))
		return nil, err
	}

	// 获取交易 Receipt
	receipt, err := i.ethclientAPI.TransactionReceipt(ctx, common.HexToHash(hash))
	if err != nil {
		ctx.Error("failed to fetch transaction receipt", zap.Error(err), zap.String("tx_hash", hash))
		return nil, err
	}

	txStatus := model.TransactionStatus_TRANSACTION_STATUS_COMPLETED
	if isPending {
		txStatus = model.TransactionStatus_TRANSACTION_STATUS_PENDING
	}

	return biz.NewTransaction(
		tx.Hash().Hex(),
		from.Hex(),
		tx.To().Hex(),
		receipt.BlockNumber.Int64(),
		tx.Time(),
	).WithReceipt(receipt).WithStatus(txStatus), nil
}
