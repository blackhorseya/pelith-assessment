//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

// ListTransactionCondition is the condition for list transaction.
type ListTransactionCondition struct {
	StartTime time.Time
	EndTime   time.Time
}

// TransactionGetter is used to get the transaction.
type TransactionGetter interface {
	ListByAddress(
		c context.Context,
		address string,
		cond ListTransactionCondition,
	) (item biz.TransactionList, total int, err error)
}

// TransactionQueryService is the service for transaction query.
type TransactionQueryService struct {
	txGetter TransactionGetter
}

// NewTransactionQueryService is used to create a new TransactionQueryService.
func NewTransactionQueryService(txGetter TransactionGetter) *TransactionQueryService {
	return &TransactionQueryService{txGetter: txGetter}
}
