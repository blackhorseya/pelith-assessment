package biz

import (
	"testing"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

func TestCampaign_AddTask(t *testing.T) {
	task := &Task{
		Task: model.Task{
			Name:        "Sample Task",
			Description: "A task for testing",
			Type:        model.TaskType_TASK_TYPE_ONBOARDING,
			Criteria:    &model.TaskCriteria{},
			Status:      model.TaskStatus_TASK_STATUS_ACTIVE,
		},
	}

	tests := []struct {
		name    string
		fields  *Campaign
		args    *Task
		wantErr bool
	}{
		{
			name: "Add task to pending campaign",
			fields: &Campaign{
				Campaign: model.Campaign{
					Status: model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
				},
			},
			args:    task,
			wantErr: false,
		},
		{
			name: "Add task to active campaign (error)",
			fields: &Campaign{
				Campaign: model.Campaign{
					Status: model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE,
				},
			},
			args:    task,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.AddTask(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCampaign_Complete(t *testing.T) {
	tests := []struct {
		name    string
		fields  *Campaign
		wantErr bool
	}{
		{
			name: "Complete active campaign",
			fields: &Campaign{
				Campaign: model.Campaign{
					Status: model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE,
				},
			},
			wantErr: false,
		},
		{
			name: "Complete pending campaign (error)",
			fields: &Campaign{
				Campaign: model.Campaign{
					Status: model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.Complete()
			if (err != nil) != tt.wantErr {
				t.Errorf("Complete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCampaign_Start(t *testing.T) {
	tests := []struct {
		name    string
		fields  *Campaign
		wantErr bool
	}{
		{
			name: "Start pending campaign",
			fields: &Campaign{
				Campaign: model.Campaign{
					Status: model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
				},
			},
			wantErr: false,
		},
		{
			name: "Start active campaign (error)",
			fields: &Campaign{
				Campaign: model.Campaign{
					Status: model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.Start()
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCampaign(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		args    string
		startAt time.Time
		wantErr bool
	}{
		{
			name:    "Valid campaign",
			args:    "Test Campaign",
			startAt: now,
			wantErr: false,
		},
		{
			name:    "Empty name (error)",
			args:    "",
			startAt: now,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCampaign(tt.args, tt.startAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCampaign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
