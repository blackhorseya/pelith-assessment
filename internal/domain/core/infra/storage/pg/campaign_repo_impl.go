package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
)

type campaignRepoImpl struct {
}

// NewCampaignRepo is used to create a new campaignRepoImpl.
func NewCampaignRepo() command.CampaignCreator {
	return &campaignRepoImpl{}
}

func (i *campaignRepoImpl) Create(c context.Context, campaign *model.Campaign) error {
	// TODO: 2024/11/20|sean|implement create campaign
	panic("implement me")
}
