package command

import (
	"context"
	"errors"

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

type backtestStrategy struct{}

func (s *backtestStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/11/24|sean|Implement the Execute method
	panic("implement me")
}

type realTimeStrategy struct{}

func (s *realTimeStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/11/24|sean|Implement the Execute method
	panic("implement me")
}

// StartCampaignHandler is the handler for starting a campaign.
type StartCampaignHandler struct {
	strategies     map[model.CampaignMode]CampaignStrategy
	campaignGetter query.CampaignGetter
}

// NewStartCampaignHandler creates a new StartCampaignHandler instance.
func NewStartCampaignHandler(campaignGetter query.CampaignGetter) *StartCampaignHandler {
	return &StartCampaignHandler{
		strategies: map[model.CampaignMode]CampaignStrategy{
			model.CampaignMode_CAMPAIGN_MODE_UNSPECIFIED: &emptyStrategy{},
			model.CampaignMode_CAMPAIGN_MODE_BACKTEST:    &backtestStrategy{},
			model.CampaignMode_CAMPAIGN_MODE_REAL_TIME:   &realTimeStrategy{},
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
