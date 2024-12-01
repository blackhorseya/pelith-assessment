package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/composite"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"go.uber.org/zap"
)

type backtestServiceImpl struct {
	bus    eventx.EventBus
	txRepo *composite.TransactionCompositeRepoImpl
}

// NewBacktestService creates a new BacktestService instance.
func NewBacktestService(txRepo *composite.TransactionCompositeRepoImpl) biz.BacktestService {
	return &backtestServiceImpl{
		txRepo: txRepo,
	}
}

func (i *backtestServiceImpl) RunBacktest(
	c context.Context,
	campaign *biz.Campaign,
	resultCh chan<- *model.Reward,
) error {
	ctx := contextx.WithContext(c)

	// 1. 獲取交易記錄
	txCh := make(chan *biz.Transaction)
	var err error
	go func() {
		err = i.txRepo.GetSwapTxByPoolAddress(ctx, campaign.PoolId, query.ListTransactionCondition{
			StartTime: campaign.StartTime.AsTime(),
			EndTime:   campaign.EndTime.AsTime(),
		}, txCh)
		if err != nil {
			ctx.Error("failed to get swapTx swapTx by pool address", zap.Error(err))
		}
		close(txCh)
	}()

	// 3. 處理交易記錄
	for swapTx := range txCh {
		reward, err2 := campaign.OnSwapExecuted(swapTx)
		if err2 != nil {
			ctx.Error("failed to handle swap executed", zap.Error(err2))
			continue
		}
		if reward == nil {
			continue
		}
		resultCh <- reward
	}

	// 4. 分配 Share Pool Task 獎勵
	rewards := campaign.GetSharePoolTaskReward()
	for _, reward := range rewards {
		resultCh <- reward
	}

	return nil
}
