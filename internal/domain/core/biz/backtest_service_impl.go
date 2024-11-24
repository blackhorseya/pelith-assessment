package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

type backtestServiceImpl struct {
}

// NewBacktestService creates a new BacktestService instance.
func NewBacktestService() biz.BacktestService {
	return &backtestServiceImpl{}
}

func (i *backtestServiceImpl) RunBacktest(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/11/24|sean|implement backtest service
	panic("implement me")
}
