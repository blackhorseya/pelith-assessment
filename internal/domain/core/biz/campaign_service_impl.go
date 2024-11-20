package biz

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
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
	// TODO: 2024/11/20|sean|implement me
	panic("implement me")
}
