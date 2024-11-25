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
	// 設定資料庫交易
	tx, err := i.rw.BeginTxx(c, nil)
	if err != nil {
		return err
	}

	// 使用 defer 確保在出現錯誤時回滾
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 插入 transaction 資料
	transactionDAO := FromBizTransactionToDAO(transaction)
	transactionQuery := `
		INSERT INTO transactions (tx_hash, block_number, timestamp, from_address, to_address)
		VALUES (:tx_hash, :block_number, :timestamp, :from_address, :to_address)`
	_, err = tx.NamedExecContext(c, transactionQuery, transactionDAO)
	if err != nil {
		return err
	}

	// 插入 swap_event 資料
	if transaction.SwapDetails != nil {
		for _, swap := range transaction.SwapDetails {
			swapEventDAO := FromModelSwapDetailToDAO(transaction.TxHash, swap)
			swapQuery := `
				INSERT INTO swap_events (
				                         tx_hash, 
				                         from_token_address, 
				                         to_token_address, 
				                         from_token_amount, 
				                         to_token_amount, 
				                         pool_address)
				VALUES (:tx_hash, :from_token_address, :to_token_address, :from_token_amount, :to_token_amount, :pool_address)`
			_, err = tx.NamedExecContext(c, swapQuery, swapEventDAO)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
