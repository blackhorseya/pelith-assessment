package biz

import (
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// Transaction is an aggregate root that represents the transaction.
type Transaction struct {
	model.Transaction
}

// NewSwapTransaction is used to create a new swap transaction.
func NewSwapTransaction(
	txHash, fromAddress, toAddress []byte,
	amount float64,
	swapDetails *model.SwapDetails,
) (*Transaction, error) {
	// TODO: 2024/11/22|sean|implement the NewTransaction
	return nil, errors.New("not implemented")
}

// TransactionList is a list of transactions.
type TransactionList []*Transaction
