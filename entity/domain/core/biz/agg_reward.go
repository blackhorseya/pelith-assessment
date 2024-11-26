package biz

import (
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// Reward is the reward aggregate.
type Reward struct {
	model.Reward
}

// NewReward is used to create a new Reward.
func NewReward(points int64, campaignID string, toAddress string) *Reward {
	return &Reward{
		Reward: model.Reward{
			Id:          "", // will be set by storage
			UserAddress: toAddress,
			CampaignId:  campaignID,
			Points:      points,
			RedeemedAt:  nil,
		},
	}
}
