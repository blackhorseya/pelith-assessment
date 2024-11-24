package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
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
	resultCh chan<- *model.BacktestResult,
) error {
	// TODO: 2024/11/24|sean|Implement the RunBacktest method
	return nil
}
