package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type backtestServiceImpl struct {
}

// NewBacktestService creates a new BacktestService instance.
func NewBacktestService() biz.BacktestService {
	return &backtestServiceImpl{}
}

func (i *backtestServiceImpl) RunBacktest(
	c context.Context,
	campaign *biz.Campaign,
	resultCh chan<- *model.BacktestResult,
) error {
	// TODO: 2024/11/24|sean|Implement the RunBacktest method
	return nil
}
