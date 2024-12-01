package pg

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// RewardRepoImpl is the implementation of RewardRepo.
type RewardRepoImpl struct {
	rw *sqlx.DB
}

// NewRewardRepo is used to create a new RewardRepo.
func NewRewardRepo(rw *sqlx.DB) (*RewardRepoImpl, error) {
	err := migrateUp(rw, "campaign")
	if err != nil {
		return nil, err
	}

	return &RewardRepoImpl{rw: rw}, nil
}

// NewRewardGetter is used to create a new RewardGetter.
func NewRewardGetter(impl *RewardRepoImpl) (query.RewardGetter, error) {
	return impl, nil
}

func (i *RewardRepoImpl) GetByAddress(c context.Context, address string) ([]*biz.Reward, error) {
	ctx := contextx.WithContext(c)

	// Set a timeout for the operation
	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// Validate input
	if address == "" {
		ctx.Error("address is required")
		return nil, errors.New("address is required")
	}

	// Prepare the query to fetch rewards by user address
	query := `
		SELECT id, user_address, campaign_id, points, redeemed_at, created_at, updated_at
		FROM rewards
		WHERE user_address = $1
	`

	// Slice to hold the fetched rewards
	var rewardDAOs []RewardDAO

	// Execute the query
	err := i.rw.SelectContext(timeout, &rewardDAOs, query, address)
	if err != nil {
		ctx.Error("failed to fetch rewards by address", zap.Error(err), zap.String("address", address))
		return nil, err
	}

	// Convert DAOs to biz.Reward
	rewards := make([]*biz.Reward, 0, len(rewardDAOs))
	for _, dao := range rewardDAOs {
		rewards = append(rewards, dao.ToAggregate())
	}

	ctx.Info("successfully fetched rewards", zap.Int("count", len(rewards)), zap.String("address", address))
	return rewards, nil
}
