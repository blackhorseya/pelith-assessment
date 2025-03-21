//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type (
	// TaskCreator is used to create a new task.
	TaskCreator interface {
		// Create is used to create a new task.
		Create(c context.Context, task *biz.Task) error
	}

	TaskUpdater interface {
		// Update is used to update a task.
		Update(c context.Context, task *biz.Task) error
	}
)

type (
	// CampaignCreator is used to create a new campaign.
	CampaignCreator interface {
		// Create is used to create a new campaign.
		Create(c context.Context, campaign *biz.Campaign) error
	}

	// CampaignUpdater is used to update the campaign.
	CampaignUpdater interface {
		DistributeReward(c context.Context, reward *model.Reward) error
		UpdateStatus(c context.Context, campaign *biz.Campaign, status model.CampaignStatus) error
	}

	// CampaignDeleter is used to delete the campaign.
	CampaignDeleter interface {
		CleanReward(c context.Context, campaignID string) error
	}
)

type (
	// UserCreator defines the interface for creating a new user.
	UserCreator interface {
		// Create persists a new user.
		Create(c context.Context, user *model.User) error
	}

	// UserUpdater defines the interface for updating user information.
	UserUpdater interface {
		// Update modifies an existing user's data.
		Update(c context.Context, user *model.User) error

		// IncrementPoints adds points to a user's total.
		IncrementPoints(c context.Context, userID string, points int64) error
	}
)

type (
	// RewardCreator is used to create a new reward.
	RewardCreator interface {
		// Create is used to create a new reward.
		Create(c context.Context, reward *biz.Reward) error
	}
)

type (
	// TransactionCreator is used to create a new transaction.
	TransactionCreator interface {
		// Create is used to create a new transaction.
		Create(c context.Context, transaction *biz.Transaction) error
	}
)
