package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// RunBacktestHandler is the handler for running a backtest.
type RunBacktestHandler struct {
	backtestService biz.BacktestService
	campaignGetter  query.CampaignGetter
}

// NewRunBacktestHandler is used to create a new RunBacktestHandler.
func NewRunBacktestHandler(
	backtestService biz.BacktestService,
	campaignGetter query.CampaignGetter,
) *RunBacktestHandler {
	return &RunBacktestHandler{
		backtestService: backtestService,
		campaignGetter:  campaignGetter,
	}
}

// Handle is used to handle the run backtest.
func (h *RunBacktestHandler) Handle(c context.Context, campaignID string, resultCh chan<- *model.BacktestResult) error {
	ctx := contextx.WithContext(c)

	campaign, err := h.campaignGetter.GetByID(c, campaignID)
	if err != nil {
		ctx.Error("failed to get campaign by ID", zap.Error(err), zap.String("campaign_id", campaignID))
		return err
	}
	if campaign == nil {
		ctx.Error("campaign not found", zap.String("campaign_id", campaignID))
		return errors.New("campaign not found")
	}

	rewards := make(chan *model.Reward)
	go func() {
		err = h.backtestService.RunBacktest(ctx, campaign, rewards)
		if err != nil {
			ctx.Error("failed to run backtest", zap.Error(err))
		}
		close(rewards)
	}()

	for reward := range rewards {
		ctx.Debug("backtest result", zap.Any("result", &reward))
		resultCh <- &model.BacktestResult{
			UserId:       reward.UserId,
			TotalSwaps:   0,
			TotalVolume:  0,
			PointsEarned: reward.Points,
			TaskProgress: nil,
		}
	}
	if err != nil {
		ctx.Error("failed to run backtest", zap.Error(err))
		return err
	}

	return nil
}
