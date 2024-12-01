package mongodb

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/repo"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
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

// NewCampaignFinder is a function to create a new campaign finder
func NewCampaignFinder(impl *CampaignRepoImpl) repo.CampaignFinder {
	return impl
}

func (i *CampaignRepoImpl) Create(c context.Context, campaign *biz.Campaign) error {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	campaign.Id = primitive.NewObjectID().Hex()
	_, err := i.coll.InsertOne(timeout, campaign)
	return err
}

func (i *CampaignRepoImpl) GetByID(c context.Context, id string) (*biz.Campaign, error) {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	var campaign *biz.Campaign
	err := i.coll.FindOne(timeout, primitive.M{"_id": id}).Decode(&campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (i *CampaignRepoImpl) List(c context.Context, cond repo.ListCampaignCondition) ([]*biz.Campaign, int64, error) {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	var campaigns []*biz.Campaign
	filter := bson.M{}

	limit, offset := defaultLimit, int64(0)
	if cond.Limit > 0 && cond.Limit <= maxLimit {
		limit = cond.Limit
	}
	if cond.Offset > 0 {
		offset = cond.Offset
	}

	opts := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := i.coll.Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("failed to find campaigns", zap.Error(err), zap.Any("filter", &filter))
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &campaigns)
	if err != nil {
		ctx.Error("failed to decode campaigns", zap.Error(err))
		return nil, 0, err
	}

	total, err := i.coll.CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("failed to count documents", zap.Error(err))
		return nil, 0, err
	}

	return campaigns, total, nil
}
