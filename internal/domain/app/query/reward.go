//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

type (
	// RewardGetter is the interface for reward getter
	RewardGetter interface {
		GetByAddress(c context.Context, address string) ([]*biz.Reward, error)
	}
)

// RewardQueryStore is the store for reward query
type RewardQueryStore struct {
	rewardGetter RewardGetter
}

// NewRewardQueryStore is the constructor for RewardQueryStore
func NewRewardQueryStore(rewardGetter RewardGetter) *RewardQueryStore {
	return &RewardQueryStore{
		rewardGetter: rewardGetter,
	}
}

// GetRewardHistoryByWalletAddress is used to get reward history by wallet address
func (s *RewardQueryStore) GetRewardHistoryByWalletAddress(c context.Context, address string) ([]*biz.Reward, error) {
	return s.rewardGetter.GetByAddress(c, address)
}
