package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
)

type userServiceImpl struct {
	campaignGetter query.CampaignGetter
	txGetter       query.TransactionGetter
}

// NewUserService creates a new UserService instance.
func NewUserService(
	campaignGetter query.CampaignGetter,
) biz.UserService {
	return &userServiceImpl{
		campaignGetter: campaignGetter,
	}
}

func (i *userServiceImpl) GetUserTaskListByAddress(c context.Context, address string) (*biz.User, error) {
	// ctx := contextx.WithContext(c)
	//
	// campaigns, total, err := i.campaignGetter.List(ctx, query.ListCampaignCondition{})
	// if err != nil {
	// 	ctx.Error("failed to list campaigns", zap.Error(err))
	// 	return nil, err
	// }
	// if total == 0 {
	// 	ctx.Error("no campaigns found")
	// 	return nil, errors.New("no campaigns found")
	// }
	//
	// user, err := biz.NewUser(address)
	// if err != nil {
	// 	ctx.Error("failed to create user", zap.Error(err))
	// 	return nil, err
	// }

	// TODO: 2024/11/26|sean|implement GetUserTaskListByAddress
	panic("implement me")
}
