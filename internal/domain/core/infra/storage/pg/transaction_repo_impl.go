package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

// TransactionRepoImpl represents the PostgreSQL implementation of the TransactionRepo.
type TransactionRepoImpl struct {
}

// NewTransactionRepoImpl creates a new TransactionRepoImpl.
func NewTransactionRepoImpl() *TransactionRepoImpl {
	return &TransactionRepoImpl{}
}

func (i *TransactionRepoImpl) Create(c context.Context, transaction *biz.Transaction) error {
	// TODO: 2024/11/24|sean|implement me
	panic("implement me")
}
