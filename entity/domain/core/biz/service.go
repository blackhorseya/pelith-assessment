//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// TaskService provides operations related to task management.
type TaskService interface {
	// EvaluateTask checks if a user has completed a specific task.
	EvaluateTask(c context.Context, userID string, taskID string) (*model.TaskResult, error)

	// CreateTask creates a new task in the system.
	CreateTask(c context.Context, task *model.Task) (string, error)
}

// CampaignService defines the domain logic for campaign management.
type CampaignService interface {
	// StartCampaign initializes a new campaign.
	StartCampaign(c context.Context, name string, startAt time.Time) (*model.Campaign, error)
}

// RewardService defines the domain logic for rewards and point allocation.
type RewardService interface {
	// AllocatePoints calculates and distributes points for a task or campaign.
	AllocatePoints(c context.Context, taskID string, poolID *string, totalPoints int64) ([]*model.PointAllocation, error)

	// RedeemReward processes a user's reward redemption.
	RedeemReward(c context.Context, userID string, campaignID string, points int64) (*model.Reward, error)
}

// UserService defines the domain logic for user management.
type UserService interface {
	// RegisterUser creates a new user in the system.
	RegisterUser(c context.Context, username string, address string) (*model.User, error)

	// UpdateUserProgress updates the progress of a specific task for a user.
	UpdateUserProgress(c context.Context, userID string, taskID string, completed bool) error
}

// BacktestService defines the domain logic for backtesting campaigns with historical data.
type BacktestService interface {
	// RunBacktest executes a backtest for a campaign within a specified time range.
	RunBacktest(c context.Context, campaignID string, startTime time.Time, endTime time.Time) error
}
