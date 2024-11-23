package biz

import (
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// Transaction is an aggregate root that represents the transaction.
type Transaction struct {
	model.Transaction
}

// TransactionList is a list of transactions.
type TransactionList []*Transaction
