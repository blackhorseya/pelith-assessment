package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// CampaignStrategy is the interface for campaign strategies.
type CampaignStrategy interface {
	Execute(c context.Context, campaign *biz.Campaign) error
}

type emptyStrategy struct{}

func (s *emptyStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	return nil
}

type backtestStrategy struct {
	backtestService biz.BacktestService
}

func (s *backtestStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	ctx := contextx.WithContext(c)

	resultCh := make(chan *model.BacktestResult)
	var err error
	go func() {
		err = s.backtestService.RunBacktest(ctx, campaign, resultCh)
		if err != nil {
			ctx.Error("failed to run backtest", zap.Error(err))
		}
		close(resultCh)
	}()

	for result := range resultCh {
		ctx.Debug("backtest result", zap.Any("result", &result))
	}

	return err
}

type realTimeStrategy struct{}

func (s *realTimeStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/11/24|sean|Implement the Execute method
	ctx := contextx.WithContext(c)
	ctx.Debug("real time strategy not implemented", zap.Any("campaign", &campaign))
	return fmt.Errorf("not implemented")
}

// StartCampaignHandler is the handler for starting a campaign.
type StartCampaignHandler struct {
	strategies     map[model.CampaignMode]CampaignStrategy
	campaignGetter query.CampaignGetter
}

// NewStartCampaignHandler creates a new StartCampaignHandler instance.
func NewStartCampaignHandler(
	campaignGetter query.CampaignGetter,
	backtestService biz.BacktestService,
) *StartCampaignHandler {
	return &StartCampaignHandler{
		strategies: map[model.CampaignMode]CampaignStrategy{
			model.CampaignMode_CAMPAIGN_MODE_UNSPECIFIED: &emptyStrategy{},
			model.CampaignMode_CAMPAIGN_MODE_BACKTEST: &backtestStrategy{
				backtestService: backtestService,
			},
			model.CampaignMode_CAMPAIGN_MODE_REAL_TIME: &realTimeStrategy{},
		},
		campaignGetter: campaignGetter,
	}
}

func (h StartCampaignHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	ctx := contextx.WithContext(c)

	cmd, ok := msg.(StartCampaignCommand)
	if !ok {
		ctx.Error("invalid command type for StartCampaignHandler", zap.Any("command", &msg))
		return "", errors.New("invalid command type for StartCampaignHandler")
	}

	campaign, err := h.campaignGetter.GetByID(ctx, cmd.ID)
	if err != nil {
		ctx.Error("failed to fetch campaign", zap.Error(err))
		return "", err
	}
	if campaign == nil {
		return "", errors.New("campaign not found")
	}

	strategy, ok := h.strategies[campaign.Mode]
	if !ok {
		ctx.Error("strategy not found", zap.Any("mode", campaign.Mode))
		return "", errors.New("strategy not found")
	}

	err = strategy.Execute(ctx, campaign)
	if err != nil {
		ctx.Error("failed to execute strategy", zap.Error(err))
		return "", err
	}

	return campaign.Id, nil
}
