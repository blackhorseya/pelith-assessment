package composite

import (
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/external/etherscan"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/storage/pg"
)

// TransactionCompositeRepoImpl represents the composite implementation of the TransactionCompositeRepo.
type TransactionCompositeRepoImpl struct {
	dbRepo  *pg.TransactionRepoImpl
	apiRepo *etherscan.TransactionRepoImpl
}

// NewTransactionCompositeRepoImpl creates a new TransactionCompositeRepoImpl instance.
func NewTransactionCompositeRepoImpl(
	dbRepo *pg.TransactionRepoImpl,
	apiRepo *etherscan.TransactionRepoImpl,
) *TransactionCompositeRepoImpl {
	return &TransactionCompositeRepoImpl{
		dbRepo:  dbRepo,
		apiRepo: apiRepo,
	}
}
