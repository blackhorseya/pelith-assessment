package biz

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"github.com/ethereum/go-ethereum/core/types"
)

// Transaction is an aggregate root that represents the transaction.
type Transaction struct {
	tx *model.Transaction

	receipt *types.Receipt

	SwapDetails []*model.SwapDetail
}

// NewTransaction is used to create a new transaction.
func NewTransaction(tx *model.Transaction, receipt *types.Receipt) *Transaction {
	return &Transaction{
		tx:          tx,
		receipt:     receipt,
		SwapDetails: tx.SwapDetails,
	}
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
