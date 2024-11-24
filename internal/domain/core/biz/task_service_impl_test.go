package biz

import (
	"context"
	"reflect"
	"testing"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

func Test_taskServiceImpl_CreateTask(t1 *testing.T) {
	type args struct {
		c           context.Context
		campaign    *biz.Campaign
		name        string
		description string
		taskType    model.TaskType
		minAmount   float64
		poolID      string
	}
	tests := []struct {
		name    string
		args    args
		want    *biz.Task
		wantErr bool
	}{
		{
			name: "Valid Onboarding Task",
			args: args{
				c: context.Background(),
				campaign: &biz.Campaign{
					Campaign: model.Campaign{
						Status: model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
					},
				},
				name:        "Onboarding Task",
				description: "Description for onboarding task",
				taskType:    model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount:   0,
				poolID:      "",
			},
			want: &biz.Task{
				Task: model.Task{
					Name:        "Onboarding Task",
					Description: "Description for onboarding task",
					Type:        model.TaskType_TASK_TYPE_ONBOARDING,
					Criteria:    &model.TaskCriteria{},
					Status:      model.TaskStatus_TASK_STATUS_ACTIVE,
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Share Pool Task",
			args: args{
				c: context.Background(),
				campaign: &biz.Campaign{
					Campaign: model.Campaign{
						Status: model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
					},
				},
				name:        "Share Pool Task",
				description: "Description for share pool task",
				taskType:    model.TaskType_TASK_TYPE_SHARE_POOL,
				minAmount:   50.0,
				poolID:      "pool123",
			},
			want: &biz.Task{
				Task: model.Task{
					Name:        "Share Pool Task",
					Description: "Description for share pool task",
					Type:        model.TaskType_TASK_TYPE_SHARE_POOL,
					Criteria: &model.TaskCriteria{
						MinTransactionAmount: 50.0,
						PoolId:               "pool123",
					},
					Status: model.TaskStatus_TASK_STATUS_ACTIVE,
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Empty Name",
			args: args{
				c: context.Background(),
				campaign: &biz.Campaign{
					Campaign: model.Campaign{
						Status: model.CampaignStatus_CAMPAIGN_STATUS_PENDING,
					},
				},
				name:        "",
				description: "Description",
				taskType:    model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount:   0,
				poolID:      "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid Task on Active Campaign",
			args: args{
				c: context.Background(),
				campaign: &biz.Campaign{
					Campaign: model.Campaign{
						Status: model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE,
					},
				},
				name:        "Task on Active Campaign",
				description: "Description",
				taskType:    model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount:   0,
				poolID:      "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := taskServiceImpl{}
			got, err := t.CreateTask(
				tt.args.c,
				tt.args.campaign,
				tt.args.name,
				tt.args.description,
				tt.args.taskType,
				tt.args.minAmount,
				tt.args.poolID,
			)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
