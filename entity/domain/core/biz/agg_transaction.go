package biz

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Transaction is an aggregate root that represents the transaction.
type Transaction struct {
	tx      *model.Transaction
	receipt *types.Receipt

	SwapDetail *model.SwapDetail
}

// NewTransaction is used to create a new transaction.
func NewTransaction(txHash, from, to string, blockNumber int64, ts time.Time) *Transaction {
	return &Transaction{
		tx: &model.Transaction{
			TxHash:      txHash,
			BlockNumber: blockNumber,
			FromAddress: from,
			ToAddress:   to,
			Amount:      0,
			Timestamp:   timestamppb.New(ts),
			TaskId:      "",
			CampaignId:  "",
			Status:      0,
			Type:        0,
			SwapDetails: nil,
		},
	}
}

// WithReceipt is used to set the receipt.
func (x *Transaction) WithReceipt(receipt *types.Receipt) *Transaction {
	x.receipt = receipt
	return x
}

// WithStatus is used to set the status.
func (x *Transaction) WithStatus(status model.TransactionStatus) *Transaction {
	x.tx.Status = status
	return x
}

// GetTransaction is used to get the transaction.
func (x *Transaction) GetTransaction() *model.Transaction {
	return x.tx
}

// GetSwapForPool is used to get the swap for the pool.
func (x *Transaction) GetSwapForPool(poolAddress common.Address, swapEventHash common.Hash) (*model.SwapDetail, error) {
	if x.receipt == nil {
		return nil, errors.New("i need receipt to get swap for pool")
	}

	var firstLog, lastLog *types.Log
	var fromTokenAddress, toTokenAddress common.Address
	var fromAmount, toAmount *big.Int

	for _, logEntry := range x.receipt.Logs {
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
			firstLog = x.receipt.Logs[0]
		}
		// Update the last valid log
		lastLog = x.receipt.Logs[len(x.receipt.Logs)-1]
	}

	if firstLog == nil || lastLog == nil {
		return nil, errors.New("no logs found")
	}

	if len(firstLog.Data) < 32 || len(lastLog.Data) < 32 {
		return nil, errors.New("log data length is insufficient")
	}

	fromTokenAddress = firstLog.Address
	toTokenAddress = lastLog.Address

	fromAmount = new(big.Int).SetBytes(firstLog.Data[:32]) // First 32 bytes represent the amount
	toAmount = new(big.Int).SetBytes(lastLog.Data[:32])    // First 32 bytes represent the amount

	x.SwapDetail = &model.SwapDetail{
		FromTokenAddress: fromTokenAddress.Hex(),
		ToTokenAddress:   toTokenAddress.Hex(),
		FromTokenAmount:  fromAmount.String(),
		ToTokenAmount:    toAmount.String(),
		PoolAddress:      poolAddress.Hex(),
	}
	return x.SwapDetail, nil
}

// IsSwapType is used to check if the transaction is swap executed.
func (x *Transaction) IsSwapType() bool {
	return x.tx.Type == model.TransactionType_TRANSACTION_TYPE_SWAP
}

// GetSwapAmountByTokenAddress is used to get the swap amount by token address.
func (x *Transaction) GetSwapAmountByTokenAddress(tokenAddress string) string {
	if x.SwapDetail == nil {
		return "0"
	}

	if strings.EqualFold(x.SwapDetail.FromTokenAddress, tokenAddress) {
		return x.SwapDetail.FromTokenAmount
	}

	if strings.EqualFold(x.SwapDetail.ToTokenAddress, tokenAddress) {
		return x.SwapDetail.ToTokenAmount
	}

	return "0"
}

// Process is used to process the transaction.
func (x *Transaction) Process(c context.Context) (eventx.DomainEvent, error) {
	if x.IsSwapType() {
		return NewSwapExecutedEvent(x.tx.Timestamp.AsTime(), SwapExecutedPayload{
			TxID: x.tx.TxHash,
		}), nil
	}

	return nil, errors.New("unsupported transaction type")
}

// TransactionList is a list of transactions.
type TransactionList []*Transaction
