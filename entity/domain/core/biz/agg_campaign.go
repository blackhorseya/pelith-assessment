package biz

import (
	"errors"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Campaign represents the aggregate root for campaigns.
type Campaign struct {
	model.Campaign

	Tasks []*Task
}

// NewCampaign creates a new Campaign aggregate.
func NewCampaign(name string, startAt time.Time) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &Campaign{
		Campaign: model.Campaign{
			Id:          "",
			Name:        name,
			Description: "",
			StartTime:   timestamppb.New(startAt),
			EndTime:     timestamppb.New(startAt.Add(4 * 7 * 24 * time.Hour)),
			Status:      model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
			Mode:        model.CampaignMode_CAMPAIGN_MODE_BACKTEST,
		},
		Tasks: make([]*Task, 0),
	}, nil
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
	c.Tasks = append(c.Tasks, task)
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
