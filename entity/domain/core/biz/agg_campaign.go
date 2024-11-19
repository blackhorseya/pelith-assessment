package biz

import (
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// Campaign represents the aggregate root for campaigns.
type Campaign struct {
	model.Campaign

	Tasks []*Task
}

// NewCampaign creates a new Campaign aggregate.
func NewCampaign(id, name, description string) (*Campaign, error) {
	if id == "" || name == "" {
		return nil, errors.New("campaign ID and name are required")
	}
	return &Campaign{
		Campaign: model.Campaign{
			Id:          id,
			Name:        name,
			Description: description,
			Status:      model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
		},
		Tasks: make([]*Task, 0),
	}, nil
}

// AddTask adds a task to the campaign.
func (c *Campaign) AddTask(task *Task) error {
	if c.Status != model.CampaignStatus_CAMPAIGN_STATUS_PENDING {
		return errors.New("tasks can only be added to pending campaigns")
	}
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
