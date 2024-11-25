package composite

import (
	"context"
	"errors"
	"sync"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/external/etherscan"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/storage/pg"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// TransactionCompositeRepoImpl represents the composite implementation of the TransactionCompositeRepo.
type TransactionCompositeRepoImpl struct {
	dbRepo  *pg.TransactionRepoImpl
	apiRepo *etherscan.TransactionRepoImpl

	locks sync.Map // 用於保存每個地址的鎖
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

func (i *TransactionCompositeRepoImpl) ListByAddress(
	c context.Context,
	address string,
	cond query.ListTransactionCondition,
) (item biz.TransactionList, total int, err error) {
	ctx := contextx.WithContext(c)

	lockKey := "lock_" + address

	lock, _ := i.locks.LoadOrStore(lockKey, &sync.Mutex{})
	mtx, ok := lock.(*sync.Mutex)
	if !ok {
		ctx.Error("failed to load lock", zap.String("lockKey", lockKey))
		return nil, 0, errors.New("failed to load lock")
	}

	mtx.Lock()
	defer func() {
		mtx.Unlock()
		i.locks.Delete(lockKey)
	}()

	item, total, err = i.dbRepo.ListByAddress(ctx, address, cond)
	if err != nil {
		ctx.Error("dbRepo.ListByAddress", zap.Error(err))
		return nil, 0, err
	}

	if total == 0 {
		item, total, err = i.apiRepo.ListByAddress(ctx, address, cond)
		if err != nil {
			ctx.Error("apiRepo.ListByAddress", zap.Error(err))
			return nil, 0, err
		}

		for _, tx := range item {
			err = i.dbRepo.Create(ctx, tx)
			if err != nil {
				ctx.Error("dbRepo.Create", zap.Error(err))
				continue
			}
		}

		return item, total, nil
	}

	return item, total, nil
}

func (i *TransactionCompositeRepoImpl) GetSwapTxByPoolAddress(
	c context.Context,
	address string,
	cond query.ListTransactionCondition,
	txCh chan<- *biz.Transaction,
) error {
	ctx := contextx.WithContext(c)

	lockKey := "lock_" + address

	lock, _ := i.locks.LoadOrStore(lockKey, &sync.Mutex{})
	mtx, ok := lock.(*sync.Mutex)
	if !ok {
		ctx.Error("failed to load lock", zap.String("lockKey", lockKey))
		return errors.New("failed to load lock")
	}

	mtx.Lock()
	defer func() {
		mtx.Unlock()
		i.locks.Delete(lockKey)
	}()

	// Step 1: 查詢本地資料庫
	item, total, err := i.dbRepo.GetLogsByAddress(ctx, address, query.GetLogsCondition{
		StartTime: cond.StartTime,
		EndTime:   cond.EndTime,
	})
	if err != nil {
		ctx.Error("dbRepo.GetLogsByAddress", zap.Error(err))
		return err
	}

	if total > 0 {
		for _, tx := range item {
			txCh <- tx
		}
		return nil
	}

	// Step 2: 從外部 API 獲取數據
	apiTxCh := make(chan *biz.Transaction)
	go func() {
		defer close(apiTxCh)
		err = i.apiRepo.GetSwapTxByPoolAddress(ctx, address, cond, apiTxCh)
		if err != nil {
			ctx.Error("apiRepo.GetSwapTxByPoolAddress", zap.Error(err))
		}
	}()

	// Step 3: 寫入資料庫並傳遞給 txCh
	for apiTx := range apiTxCh {
		select {
		case txCh <- apiTx: // 傳遞數據到調用方的 channel
		case <-ctx.Done(): // 如果 context 被取消，停止操作
			ctx.Error("context cancelled while sending transaction", zap.Error(ctx.Err()))
			return ctx.Err()
		}

		// 儲存數據到資料庫
		saveErr := i.dbRepo.Create(ctx, apiTx)
		if saveErr != nil {
			ctx.Error("dbRepo.Create", zap.Error(saveErr))
			// 儲存失敗不會中斷整體邏輯，繼續處理其他交易
			continue
		}
	}

	if err != nil {
		ctx.Error("apiRepo.GetSwapTxByPoolAddress", zap.Error(err))
		return err
	}

	return nil
}
