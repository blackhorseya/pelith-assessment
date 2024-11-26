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
	campaignUpdater CampaignUpdater
}

// NewRunBacktestHandler is used to create a new RunBacktestHandler.
func NewRunBacktestHandler(
	backtestService biz.BacktestService,
	campaignGetter query.CampaignGetter,
	campaignUpdater CampaignUpdater,
) *RunBacktestHandler {
	return &RunBacktestHandler{
		backtestService: backtestService,
		campaignGetter:  campaignGetter,
		campaignUpdater: campaignUpdater,
	}
}

// Handle is used to handle the run backtest.
func (h *RunBacktestHandler) Handle(c context.Context, campaignID string, resultCh chan<- *model.Reward) error {
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
		err = h.campaignUpdater.DistributeReward(ctx, reward)
		if err != nil {
			ctx.Error("failed to distribute reward", zap.Error(err))
			continue
		}
		resultCh <- reward
	}
	if err != nil {
		ctx.Error("failed to run backtest", zap.Error(err))
		return err
	}

	return nil
}
