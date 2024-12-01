package mongodb

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/repo"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

// CampaignRepoImpl is a struct to define the campaign repository implementation
type CampaignRepoImpl struct {
	coll *mongo.Collection
}

// NewCampaignRepoImpl is a function to create a new campaign repository implementation
func NewCampaignRepoImpl(rw *mongo.Client) *CampaignRepoImpl {
	return &CampaignRepoImpl{
		coll: rw.Database(dbName).Collection("campaigns"),
	}
}

// NewCampaignCreator is a function to create a new campaign creator
func NewCampaignCreator(impl *CampaignRepoImpl) repo.CampaignCreator {
	return impl
}

func (i *CampaignRepoImpl) Create(c context.Context, campaign *biz.Campaign) error {
	ctx := contextx.WithContext(c)

	_, err := i.coll.InsertOne(ctx, campaign)
	return err
}
