package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	query2 "github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

type userServiceImpl struct {
	campaignGetter query2.CampaignGetter
	txRepo         query2.TransactionRepo
}

// NewUserService creates a new UserService instance.
func NewUserService(
	campaignGetter query2.CampaignGetter,
	txRepo query2.TransactionRepo,
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

	txCh := make(chan *biz.Transaction)
	go func() {
		err = i.txRepo.GetSwapTxByUserAddressAndPoolAddress(
			ctx,
			address,
			campaign.PoolId,
			query2.ListTransactionCondition{
				PoolAddress: campaign.PoolId,
				StartTime:   campaign.StartTime.AsTime(),
				EndTime:     campaign.EndTime.AsTime(),
			},
			txCh,
		)
		if err != nil {
			ctx.Error("failed to fetch swap transactions", zap.Error(err))
		}
		close(txCh)
	}()
	if err != nil {
		ctx.Error("failed to fetch swap transactions", zap.Error(err))
		return nil, err
	}

	for tx := range txCh {
		err = user.OnSwapExecuted(ctx, tx)
		if err != nil {
			ctx.Error("failed to process swap transaction", zap.Error(err))
			return nil, err
		}
	}

	return user, nil
}
