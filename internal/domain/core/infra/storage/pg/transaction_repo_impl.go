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
	// 統一管理查詢參數
	params := map[string]interface{}{
		"contract_address": contractAddress,
		"start_time":       cond.StartTime,
		"end_time":         cond.EndTime,
	}

	// 查詢符合條件的總筆數
	countQuery := `
		SELECT COUNT(*)
		FROM swap_events se
		JOIN transactions t ON se.tx_hash = t.tx_hash
		WHERE se.pool_address = :contract_address
		  AND t.timestamp BETWEEN :start_time AND :end_time`
	err = i.rw.GetContext(c, &total, countQuery, params)
	if err != nil {
		return nil, 0, err
	}

	// 如果總筆數為 0，直接返回
	if total == 0 {
		return biz.TransactionList{}, 0, nil
	}

	// 查詢符合條件的交易資料
	logsQuery := `
		SELECT se.tx_hash, se.from_token_address, se.to_token_address, 
		       se.from_token_amount, se.to_token_amount, se.pool_address,
		       t.block_number, t.timestamp, t.from_address, t.to_address
		FROM swap_events se
		JOIN transactions t ON se.tx_hash = t.tx_hash
		WHERE se.pool_address = :contract_address
		  AND t.timestamp BETWEEN :start_time AND :end_time
		ORDER BY t.timestamp DESC`
	var rows []struct {
		TransactionDAO
		SwapEventDAO
	}
	err = i.rw.SelectContext(c, &rows, logsQuery, params)
	if err != nil {
		return nil, 0, err
	}

	// 將查詢結果轉換為 biz.TransactionList
	item = biz.TransactionList{}
	for _, row := range rows {
		swapDetail := row.SwapEventDAO.ToModel()
		transaction := row.TransactionDAO.ToBizModel()
		transaction.SwapDetails = append(transaction.SwapDetails, swapDetail)
		item = append(item, transaction)
	}

	return item, total, nil
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
