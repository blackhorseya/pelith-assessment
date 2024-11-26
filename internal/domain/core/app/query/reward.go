//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type (
	// RewardGetter is the interface for reward getter
	RewardGetter interface {
		GetByAddress(c context.Context, address string) ([]*biz.Reward, error)
	}
)

// RewardQueryStore is the store for reward query
type RewardQueryStore struct {
}

// NewRewardQueryStore is the constructor for RewardQueryStore
func NewRewardQueryStore() *RewardQueryStore {
	return &RewardQueryStore{}
}

// GetRewardHistoryByWalletAddress is used to get reward history by wallet address
func (s *RewardQueryStore) GetRewardHistoryByWalletAddress(c context.Context, address string) ([]*model.Reward, error) {
	// TODO: 2024/11/26|sean|implement the handler
	return nil, errors.New("not implemented")
}
