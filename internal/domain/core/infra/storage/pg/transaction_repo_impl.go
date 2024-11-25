package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
)

// TransactionRepoImpl represents the PostgreSQL implementation of the TransactionRepo.
type TransactionRepoImpl struct {
}

// NewTransactionRepoImpl creates a new TransactionRepoImpl.
func NewTransactionRepoImpl() *TransactionRepoImpl {
	return &TransactionRepoImpl{}
}

func (i *TransactionRepoImpl) ListByAddress(
	c context.Context,
	address string,
	cond query.ListTransactionCondition,
) (item biz.TransactionList, total int, err error) {
	// TODO: 2024/11/25|sean|implement me
	panic("implement me")
}

func (i *TransactionRepoImpl) GetLogsByAddress(
	c context.Context,
	contractAddress string,
	cond query.GetLogsCondition,
) (item biz.TransactionList, total int, err error) {
	// TODO: 2024/11/25|sean|implement me
	panic("implement me")
}

func (i *TransactionRepoImpl) Create(c context.Context, transaction *biz.Transaction) error {
	// TODO: 2024/11/24|sean|implement me
	panic("implement me")
}
