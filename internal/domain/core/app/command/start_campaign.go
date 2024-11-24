package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
)

// CampaignStrategy is the interface for campaign strategies.
type CampaignStrategy interface {
	Execute(c context.Context, campaign *biz.Campaign) error
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
	strategies map[model.CampaignMode]CampaignStrategy
}

// NewStartCampaignHandler creates a new StartCampaignHandler instance.
func NewStartCampaignHandler() *StartCampaignHandler {
	return &StartCampaignHandler{
		strategies: map[model.CampaignMode]CampaignStrategy{
			model.CampaignMode_CAMPAIGN_MODE_BACKTEST:  &backtestStrategy{},
			model.CampaignMode_CAMPAIGN_MODE_REAL_TIME: &realTimeStrategy{},
		},
	}
}

func (h StartCampaignHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	// TODO: 2024/11/24|sean|Implement the Handle method
	panic("implement me")
}
