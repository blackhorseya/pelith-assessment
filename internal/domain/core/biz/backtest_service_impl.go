package biz

import (
	"context"
	"strconv"

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
	// send test result to channel
	for idx := 0; idx < 10; idx++ {
		resultCh <- &model.BacktestResult{
			UserId:       "test" + strconv.Itoa(idx),
			TotalSwaps:   0,
			TotalVolume:  0,
			PointsEarned: 0,
			TaskProgress: nil,
		}
	}

	return nil
}
