package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// TransactionRepoImpl represents the PostgreSQL implementation of the TransactionRepo.
type TransactionRepoImpl struct {
	rw *sqlx.DB
}

func (i *TransactionRepoImpl) GetByHash(c context.Context, hash string) (item *biz.Transaction, err error) {
	// 設定上下文，支援超時與日誌記錄
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 查詢交易資料
	stmt := `
		SELECT t.tx_hash, t.block_number, t.timestamp, t.from_address, t.to_address,
		       se.id, se.from_token_address, se.to_token_address, se.from_token_amount, 
		       se.to_token_amount, se.pool_address
		FROM transactions t
		LEFT JOIN swap_events se ON t.tx_hash = se.tx_hash
		WHERE t.tx_hash = $1`
	var row struct {
		TransactionDAO
		SwapEventDAO
	}
	err = i.rw.GetContext(timeout, &row, stmt, hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		ctx.Error("failed to fetch transaction", zap.Error(err))
		return nil, err
	}

	// 將 DAO 轉為業務模型
	transaction := row.TransactionDAO.ToBizModel()
	if row.SwapEventDAO.ID != 0 { // 確認有關聯的 SwapEvent
		transaction.SwapDetail = row.SwapEventDAO.ToModel()
	}

	return transaction, nil
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
	// 設定上下文，支援超時與日誌記錄
	ctx := contextx.WithContext(c)
	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 查詢參數設置
	params := []interface{}{address, address, cond.StartTime, cond.EndTime}

	// 建立 WHERE 條件動態查詢
	baseCondition := `
		(t.from_address = $1 OR t.to_address = $2)
		AND t.timestamp BETWEEN $3 AND $4`
	if cond.PoolAddress != "" {
		baseCondition += " AND se.pool_address = $5"
		params = append(params, cond.PoolAddress)
	}

	// 查詢符合條件的總筆數
	countQuery := `
		SELECT COUNT(*)
		FROM transactions t
		LEFT JOIN swap_events se ON t.tx_hash = se.tx_hash
		WHERE ` + baseCondition
	err = i.rw.GetContext(timeout, &total, countQuery, params...)
	if err != nil {
		ctx.Error("failed to count transactions", zap.Error(err))
		return nil, 0, err
	}

	// 如果總筆數為 0，直接返回
	if total == 0 {
		return biz.TransactionList{}, 0, nil
	}

	// 查詢符合條件的交易資料
	stmt := `
		SELECT t.tx_hash, t.block_number, t.timestamp, t.from_address, t.to_address,
		       se.from_token_address, se.to_token_address, se.from_token_amount, 
		       se.to_token_amount, se.pool_address
		FROM transactions t
		LEFT JOIN swap_events se ON t.tx_hash = se.tx_hash
		WHERE ` + baseCondition + `
		ORDER BY t.timestamp DESC`
	var rows []struct {
		TransactionDAO
		SwapEventDAO
	}
	err = i.rw.SelectContext(timeout, &rows, stmt, params...)
	if err != nil {
		ctx.Error("failed to fetch transactions", zap.Error(err))
		return nil, 0, err
	}

	// 將 DAO 轉為業務模型
	var transactions biz.TransactionList
	for _, row := range rows {
		transaction := row.TransactionDAO.ToBizModel()
		if row.SwapEventDAO.ID != 0 { // 確認有關聯的 SwapEvent
			transaction.SwapDetail = row.SwapEventDAO.ToModel()
		}
		transactions = append(transactions, transaction)
	}

	return transactions, total, nil
}

func (i *TransactionRepoImpl) GetLogsByAddress(
	c context.Context,
	contractAddress string,
	cond query.GetLogsCondition,
) (item biz.TransactionList, total int, err error) {
	// 使用 contextx.WithContext 來設置上下文
	ctx := contextx.WithContext(c)

	// 設置超時機制
	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 初始化查詢語句和參數
	baseCondition := `
		se.pool_address = $1
		AND t.timestamp BETWEEN $2 AND $3`
	params := []interface{}{contractAddress, cond.StartTime, cond.EndTime}

	// 查詢符合條件的總筆數
	countQuery := `
		SELECT COUNT(*)
		FROM swap_events se
		JOIN transactions t ON se.tx_hash = t.tx_hash
		WHERE ` + baseCondition

	// 執行查詢總筆數
	err = i.rw.GetContext(timeout, &total, countQuery, params...)
	if err != nil {
		ctx.Error("failed to count swap events", zap.Error(err))
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
		WHERE ` + baseCondition + `
		ORDER BY t.timestamp DESC`

	var rows []struct {
		TransactionDAO
		SwapEventDAO
	}
	err = i.rw.SelectContext(timeout, &rows, logsQuery, params...)
	if err != nil {
		ctx.Error("failed to fetch swap events", zap.Error(err))
		return nil, 0, err
	}

	// 將查詢結果轉換為 biz.TransactionList
	item = biz.TransactionList{}
	for _, row := range rows {
		transaction := row.TransactionDAO.ToBizModel()
		transaction.SwapDetail = row.SwapEventDAO.ToModel()
		item = append(item, transaction)
	}

	// Log 最終查詢結果
	ctx.Info("fetched swap events successfully",
		zap.Int("total", total),
		zap.Int("fetched_rows", len(rows)),
	)

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
	if transaction.SwapDetail != nil {
		swapEventDAO := FromModelSwapDetailToDAO(transaction.GetTransaction().TxHash, transaction.SwapDetail)
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

	return nil
}
