package mongodb

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/repo"
)

// CampaignRepoImpl is a struct to define the campaign repository implementation
type CampaignRepoImpl struct {
}

// NewCampaignRepoImpl is a function to create a new campaign repository implementation
func NewCampaignRepoImpl() *CampaignRepoImpl {
	return &CampaignRepoImpl{}
}

// NewCampaignCreator is a function to create a new campaign creator
func NewCampaignCreator(impl *CampaignRepoImpl) repo.CampaignCreator {
	return impl
}

func (i *CampaignRepoImpl) Create(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/12/1|sean|implement me
	panic("implement me")
}
