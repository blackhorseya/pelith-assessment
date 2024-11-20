package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
)

// CampaignRepoImpl is the campaign repository implementation.
type CampaignRepoImpl struct {
}

// NewCampaignRepo is used to create a new CampaignRepoImpl.
func NewCampaignRepo() command.CampaignCreator {
	return &CampaignRepoImpl{}
}

func (i *CampaignRepoImpl) Create(c context.Context, campaign *model.Campaign) error {
	// TODO: 2024/11/20|sean|implement create campaign
	panic("implement me")
}
