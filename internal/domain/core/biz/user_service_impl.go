package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

type userServiceImpl struct {
	campaignGetter query.CampaignGetter
	txRepo         query.TransactionRepo
}

// NewUserService creates a new UserService instance.
func NewUserService(
	campaignGetter query.CampaignGetter,
	txRepo query.TransactionRepo,
) biz.UserService {
	return &userServiceImpl{
		campaignGetter: campaignGetter,
		txRepo:         txRepo,
	}
}

func (i *userServiceImpl) GetUserTaskListByAddress(c context.Context, address, campaignID string) (*biz.User, error) {
	ctx := contextx.WithContext(c)

	campaign, err := i.campaignGetter.GetByID(ctx, campaignID)
	if err != nil || campaign == nil {
		ctx.Error("failed to fetch campaign", zap.Error(err), zap.String("campaign_id", campaignID))
		return nil, err
	}

	user, err := biz.NewUser(address)
	if err != nil {
		ctx.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	user.Tasks = campaign.Tasks()

	// TODO: 2024/11/26|sean|implement GetUserTaskListByAddress
	return user, nil
}
