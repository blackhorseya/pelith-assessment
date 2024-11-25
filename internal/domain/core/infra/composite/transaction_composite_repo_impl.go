package composite

import (
	"context"
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
	mtx := lock.(*sync.Mutex)

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

func (i *TransactionCompositeRepoImpl) GetLogsByAddress(
	c context.Context,
	contractAddress string,
	cond query.GetLogsCondition,
) (item biz.TransactionList, total int, err error) {
	ctx := contextx.WithContext(c)

	// 鎖的 Key
	lockKey := "lock_" + contractAddress

	// 加載或創建新的 Mutex
	mutex, _ := i.locks.LoadOrStore(lockKey, &sync.Mutex{})
	mtx := mutex.(*sync.Mutex)

	// 加鎖
	mtx.Lock()
	defer func() {
		mtx.Unlock()
		i.locks.Delete(lockKey) // 解鎖後刪除，避免內存泄漏
	}()

	// Step 1: 直接從外部 API 獲取數據
	apiData, apiTotal, apiErr := i.apiRepo.GetLogsByAddress(ctx, contractAddress, cond)
	if apiErr != nil {
		ctx.Error("apiRepo.GetLogsByAddress", zap.Error(apiErr))
		return nil, 0, apiErr
	}

	// Step 2: 保存數據到資料庫
	for _, tx := range apiData {
		saveErr := i.dbRepo.Create(ctx, tx)
		if saveErr != nil {
			// 日誌記錄，但不影響主邏輯
			ctx.Error("dbRepo.Create", zap.Error(saveErr))
			continue
		}
	}

	// Step 3: 返回數據
	return apiData, apiTotal, nil
}
