package biz

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
)

// Transaction is an aggregate root that represents the transaction.
type Transaction struct {
	model.Transaction
}

// IsSwapExecuted is used to check if the transaction is swap executed.
func (tx *Transaction) IsSwapExecuted() bool {
	return tx.Type == model.TransactionType_TRANSACTION_TYPE_SWAP
}

// Process is used to process the transaction.
func (tx *Transaction) Process(c context.Context) (eventx.DomainEvent, error) {
	if tx.IsSwapExecuted() {
		return NewSwapExecutedEvent(tx.Timestamp.AsTime(), SwapExecutedPayload{
			TxID: tx.TxHash,
		}), nil
	}

	return nil, errors.New("unsupported transaction type")
}

// TransactionList is a list of transactions.
type TransactionList []*Transaction
