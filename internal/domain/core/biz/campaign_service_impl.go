package biz

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

type campaignServiceImpl struct {
}

// NewCampaignService creates a new CampaignService instance.
func NewCampaignService() biz.CampaignService {
	return &campaignServiceImpl{}
}

func (i *campaignServiceImpl) CreateCampaign(c context.Context, name string, startAt time.Time) (*biz.Campaign, error) {
	// TODO: 2024/11/24|sean|implement me
	panic("implement me")
}

func (i *campaignServiceImpl) StartCampaign(
	c context.Context,
	name string,
	startAt time.Time,
	tasks []*biz.Task,
) (*biz.Campaign, error) {
	ctx := contextx.WithContext(c)

	campaign, err := biz.NewCampaign(name, startAt)
	if err != nil {
		ctx.Error("failed to create campaign", zap.Error(err))
		return nil, err
	}

	for _, task := range tasks {
		err = campaign.AddTask(task)
		if err != nil {
			ctx.Error("failed to add task to campaign", zap.Error(err))
			return nil, err
		}
	}

	// TODO: 2024/11/20|sean|maybe call start campaign.Start() here

	return campaign, nil
}
