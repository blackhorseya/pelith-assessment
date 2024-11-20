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

func (i *campaignServiceImpl) StartCampaign(
	c context.Context,
	name string,
	startAt time.Time,
) (*model.Campaign, error) {
	ctx := contextx.WithContext(c)

	campaign, err := biz.NewCampaign(name, startAt)
	if err != nil {
		ctx.Error("failed to create campaign", zap.Error(err))
		return nil, err
	}

	// TODO: 2024/11/20|sean|maybe you need add tasks to campaign

	return &campaign.Campaign, nil
}
