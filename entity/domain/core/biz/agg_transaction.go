package biz

import (
	"context"
	"errors"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Transaction is an aggregate root that represents the transaction.
type Transaction struct {
	tx *model.Transaction

	receipt *types.Receipt

	SwapDetails []*model.SwapDetail
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

// IsSwapType is used to check if the transaction is swap executed.
func (x *Transaction) IsSwapType() bool {
	return x.tx.Type == model.TransactionType_TRANSACTION_TYPE_SWAP
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
