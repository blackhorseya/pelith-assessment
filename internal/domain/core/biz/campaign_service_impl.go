package biz

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

type campaignServiceImpl struct {
}

// NewCampaignService creates a new CampaignService instance.
func NewCampaignService() biz.CampaignService {
	return &campaignServiceImpl{}
}

func (i *campaignServiceImpl) CreateCampaign(
	c context.Context,
	name string,
	startAt time.Time,
	mode model.CampaignMode,
	targetPool string,
	minAmount float64,
) (*biz.Campaign, error) {
	ctx := contextx.WithContext(c)

	campaign, err := biz.NewCampaign(name, startAt)
	if err != nil {
		ctx.Error("failed to create campaign", zap.Error(err))
		return nil, err
	}
	campaign.Mode = mode

	taskOfOnboarding, err := biz.NewTaskOfOnboarding("onboarding", "", minAmount, targetPool)
	if err != nil {
		ctx.Error("failed to create onboarding task", zap.Error(err))
		return nil, err
	}

	err = campaign.AddTask(taskOfOnboarding)
	if err != nil {
		ctx.Error("failed to add onboarding task to campaign", zap.Error(err))
		return nil, err
	}

	taskOfSharePool, err := biz.NewTaskOfSharePool("share pool", "", targetPool)
	if err != nil {
		ctx.Error("failed to create share pool task", zap.Error(err))
		return nil, err
	}

	err = campaign.AddTask(taskOfSharePool)
	if err != nil {
		ctx.Error("failed to add share pool task to campaign", zap.Error(err))
		return nil, err
	}

	return campaign, nil
}

func (i *campaignServiceImpl) StartCampaign(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/11/24|sean|Implement the StartCampaign method
	panic("implement me")
}
