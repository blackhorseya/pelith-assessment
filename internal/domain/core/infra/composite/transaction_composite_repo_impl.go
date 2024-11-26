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

// NewTransactionRepoImpl creates a new TransactionRepoImpl.
func NewTransactionRepoImpl(impl *TransactionCompositeRepoImpl) query.TransactionRepo {
	return impl
}

func (i *TransactionCompositeRepoImpl) GetSwapTxByUserAddressAndPoolAddress(
	c context.Context,
	address, poolAddress string,
	cond query.ListTransactionCondition,
	txCh chan<- *biz.Transaction,
) error {
	// TODO: 2024/11/26|sean|implement GetSwapTxByUserAddressAndPoolAddress
	return errors.New("implement GetSwapTxByUserAddressAndPoolAddress")
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

	// Step 1: 從外部 API 獲取交易哈希列表
	txHashes, _, err := i.apiRepo.GetLogsByAddress(ctx, address, query.GetLogsCondition{
		StartTime: cond.StartTime,
		EndTime:   cond.EndTime,
	})
	if err != nil {
		ctx.Error("apiRepo.GetLogsByAddress", zap.Error(err))
		return err
	}

	// Step 2: 遍歷交易哈希並檢查資料庫
	for _, txHash := range txHashes {
		// 查詢資料庫是否存在此交易
		tx, _ := i.dbRepo.GetByHash(ctx, txHash)
		if tx != nil {
			// 如果資料庫中存在，直接傳遞
			txCh <- tx
			continue
		}

		// Step 3: 如果資料庫中不存在，通過外部 API 補齊數據
		tx, err = i.apiRepo.GetByHashWithPool(ctx, txHash, address)
		if err != nil {
			ctx.Error("apiRepo.GetTxByHash", zap.String("txHash", txHash), zap.Error(err))
			continue // 忽略失敗的哈希，處理其他
		}

		// 將交易詳細信息傳遞給調用方
		txCh <- tx

		// 儲存交易到資料庫
		saveErr := i.dbRepo.Create(ctx, tx)
		if saveErr != nil {
			ctx.Error("dbRepo.Create", zap.Error(saveErr))
			// 儲存失敗不影響主邏輯，繼續處理其他
		}
	}

	return nil
}
