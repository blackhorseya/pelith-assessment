package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/jmoiron/sqlx"
)

// TransactionRepoImpl represents the PostgreSQL implementation of the TransactionRepo.
type TransactionRepoImpl struct {
	rw *sqlx.DB
}

// NewTransactionRepoImpl creates a new TransactionRepoImpl.
func NewTransactionRepoImpl(rw *sqlx.DB) (*TransactionRepoImpl, error) {
	err := migrateUp(rw, "transaction")
	if err != nil {
		return nil, err
	}

	return &TransactionRepoImpl{
		rw: rw,
	}, nil
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
	dao := &TransactionDAO{
		TxHash:      transaction.TxHash,
		BlockNumber: transaction.BlockNumber,
		Timestamp:   transaction.Timestamp.AsTime(),
		FromAddress: transaction.FromAddress,
		ToAddress:   transaction.ToAddress,
		Value:       transaction.Amount,
		Status:      true,
	}

	stmt := `
		INSERT INTO transactions (
			tx_hash, block_number, timestamp, from_address, to_address, value, gas_used, gas_price, status
		) VALUES (
			:tx_hash, :block_number, :timestamp, :from_address, :to_address, :value, :gas_used, :gas_price, :status
		)
	`
	_, err := i.rw.NamedExecContext(c, stmt, dao)
	return err
}
