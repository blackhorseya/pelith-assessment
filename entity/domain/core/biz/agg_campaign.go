package biz

import (
	"errors"
	"strconv"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Campaign represents the aggregate root for campaigns.
type Campaign struct {
	model.Campaign `bson:",inline"`

	tasks   []*Task
	rewards []*model.Reward

	transactionList      TransactionList
	userSwapVolume       map[string]float64
	userOnboardingReward map[string]bool
	totalSwapVolume      float64
}

// NewCampaign creates a new Campaign aggregate.
func NewCampaign(name string, startAt time.Time, poolID string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if startAt.IsZero() {
		return nil, errors.New("start time cannot be empty")
	}

	return &Campaign{
		Campaign: model.Campaign{
			Id:          "",
			Name:        name,
			Description: "",
			StartTime:   timestamppb.New(startAt),
			EndTime:     timestamppb.New(startAt.Add(4 * 7 * 24 * time.Hour)),
			Tasks:       nil,
			Mode:        model.CampaignMode_CAMPAIGN_MODE_BACKTEST,
			Status:      model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
			PoolId:      poolID,
		},
		userSwapVolume:       make(map[string]float64),
		userOnboardingReward: make(map[string]bool),
	}, nil
}

// Tasks returns the tasks associated with the campaign.
func (c *Campaign) Tasks() []*Task {
	return c.tasks
}

// GetTaskByType returns the task of the specified type.
func (c *Campaign) GetTaskByType(taskType model.TaskType) *Task {
	for _, task := range c.Tasks() {
		if task.Type == taskType {
			return task
		}
	}

	return nil
}

// AddTask adds a task to the campaign.
func (c *Campaign) AddTask(task *Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}

	task.CampaignID = c.Id
	c.tasks = append(c.tasks, task)
	return nil
}

// Start marks the campaign as active.
func (c *Campaign) Start() error {
	c.Status = model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE
	return nil
}

// Complete marks the campaign as completed.
func (c *Campaign) Complete() error {
	if c.Status != model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE {
		return errors.New("only active campaigns can be completed")
	}
	c.Status = model.CampaignStatus_CAMPAIGN_STATUS_COMPLETED
	return nil
}

// OnSwapExecuted handles the swap executed event.
func (c *Campaign) OnSwapExecuted(tx *Transaction) (*model.Reward, error) {
	c.transactionList = append(c.transactionList, tx)

	amount, err := strconv.ParseFloat(tx.GetSwapAmountByTokenAddress(c.PoolId), 64)
	if err != nil {
		return nil, err
	}
	c.userSwapVolume[tx.GetTransaction().FromAddress] += amount
	c.totalSwapVolume += amount

	var reward *model.Reward
	if !c.userOnboardingReward[tx.GetTransaction().FromAddress] {
		totalAmount := c.userSwapVolume[tx.GetTransaction().FromAddress]
		if c.HasCompletedOnboardingTask(totalAmount) {
			c.userOnboardingReward[tx.GetTransaction().FromAddress] = true
			reward = &model.Reward{
				Id:          "", // generate unique ID from repository
				UserAddress: tx.GetTransaction().FromAddress,
				CampaignId:  c.Id,
				Points:      100, // 固定獎勵點數
			}

			c.rewards = append(c.rewards, reward)
		}
	}

	return reward, nil
}

// HasCompletedOnboardingTask checks if the user has completed the onboarding task.
func (c *Campaign) HasCompletedOnboardingTask(volume float64) bool {
	if c == nil {
		return false
	}

	for _, task := range c.Tasks() {
		if task.Type == model.TaskType_TASK_TYPE_ONBOARDING && volume >= task.Criteria.MinTransactionAmount {
			return true
		}
	}

	return false
}

// GetSharePoolTaskReward returns the rewards for the share pool task.
func (c *Campaign) GetSharePoolTaskReward() []*model.Reward {
	var rewards []*model.Reward
	for user, volume := range c.userSwapVolume {
		if !c.HasCompletedOnboardingTask(volume) {
			continue
		}

		points := int64((volume / c.totalSwapVolume) * 10000)

		reward := &model.Reward{
			Id:          "", // generate unique ID from repository
			UserAddress: user,
			CampaignId:  c.Id,
			Points:      points,
		}

		rewards = append(rewards, reward)
	}

	return rewards
}
