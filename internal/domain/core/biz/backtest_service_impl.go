package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"go.uber.org/zap"
)

type backtestServiceImpl struct {
	bus      eventx.EventBus
	txGetter query.TransactionGetter
}

// NewBacktestService creates a new BacktestService instance.
func NewBacktestService(txGetter query.TransactionGetter) biz.BacktestService {
	return &backtestServiceImpl{
		txGetter: txGetter,
	}
}

func (i *backtestServiceImpl) RunBacktest(
	c context.Context,
	campaign *biz.Campaign,
	resultCh chan<- *model.Reward,
) error {
	ctx := contextx.WithContext(c)

	transactionList, _, err := i.txGetter.GetLogsByAddress(c, campaign.PoolId, query.GetLogsCondition{
		StartTime: campaign.StartTime.AsTime(),
		EndTime:   campaign.EndTime.AsTime(),
	})
	if err != nil {
		ctx.Error("failed to get logs by address", zap.Error(err))
		return err
	}

	// TODO: 2024/11/24|sean|process the transaction list
	for _, tx := range transactionList {
		if tx.IsSwapExecuted() {
			event, err2 := campaign.OnSwapExecuted(tx)
			if err2 != nil {
				ctx.Error("failed to handle swap executed event", zap.Error(err2))
				continue
			}

			if event != nil {
				ctx.Debug("reward event", zap.Any("event", &event))
			}
		}
	}

	return nil
}
