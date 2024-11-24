package biz

import (
	"errors"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Campaign represents the aggregate root for campaigns.
type Campaign struct {
	model.Campaign

	tasks  []*Task
	report *model.BacktestReport
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
		report: new(model.BacktestReport),
	}, nil
}

// Tasks returns the tasks associated with the campaign.
func (c *Campaign) Tasks() []*Task {
	return c.tasks
}

// Report returns the backtest report for the campaign.
func (c *Campaign) Report() *model.BacktestReport {
	return c.report
}

// AddTask adds a task to the campaign.
func (c *Campaign) AddTask(task *Task) error {
	if c.Status != model.CampaignStatus_CAMPAIGN_STATUS_PENDING {
		return errors.New("tasks can only be added to pending campaigns")
	}

	if task == nil {
		return errors.New("task cannot be nil")
	}

	task.CampaignID = c.Id
	c.tasks = append(c.tasks, task)
	return nil
}

// Start marks the campaign as active.
func (c *Campaign) Start() error {
	if c.Status != model.CampaignStatus_CAMPAIGN_STATUS_PENDING {
		return errors.New("only pending campaigns can be started")
	}
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
func (c *Campaign) OnSwapExecuted(tx *Transaction) (eventx.DomainEvent, error) {
	// TODO: 2024/11/24|sean|Add swap logic
	return nil, errors.New("implement OnSwapExecuted")
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
