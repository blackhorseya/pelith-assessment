package biz

import (
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// User is an aggregate root that represents the user.
type User struct {
	model.User

	Rewards []*model.Reward
}

// NewUser creates a new User aggregate.
func NewUser(id, name, address string) (*User, error) {
	if id == "" || name == "" {
		return nil, errors.New("user ID and name are required")
	}

	return &User{
		User: model.User{
			Id:             id,
			Name:           name,
			Address:        address,
			TaskProgress:   make(map[string]bool),
			Points:         0,
			TransactionIds: make([]string, 0),
		},
		Rewards: make([]*model.Reward, 0),
	}, nil
}

// CompleteTask updates the user's progress and awards points for a completed task.
func (u *User) CompleteTask(taskID string, points int64) error {
	if u.TaskProgress[taskID] {
		return errors.New("task already completed")
	}
	u.TaskProgress[taskID] = true
	u.Points += points
	return nil
}

// RedeemReward redeems a reward using the user's points.
func (u *User) RedeemReward(reward *model.Reward) error {
	if reward.Points > u.Points {
		return errors.New("insufficient points to redeem the reward")
	}
	u.Points -= reward.Points
	u.Rewards = append(u.Rewards, reward)
	return nil
}
