package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

func TestNewTask(t *testing.T) {
	type args struct {
		name        string
		description string
		taskType    model.TaskType
		minAmount   float64
		poolID      string
	}
	tests := []struct {
		name    string
		args    args
		want    *Task
		wantErr bool
	}{
		{
			name: "Valid task creation",
			args: args{
				name:        "Test Task",
				description: "This is a test task.",
				taskType:    model.TaskType_TASK_TYPE_SHARE_POOL,
				minAmount:   100.0,
				poolID:      "pool-123",
			},
			want: &Task{
				Task: model.Task{
					Name:        "Test Task",
					Description: "This is a test task.",
					Type:        model.TaskType_TASK_TYPE_SHARE_POOL,
					Criteria: &model.TaskCriteria{
						MinTransactionAmount: 100.0,
						PoolId:               "pool-123",
					},
					Status: model.TaskStatus_TASK_STATUS_ACTIVE,
				},
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			args: args{
				name:        "",
				description: "This task has no name.",
				taskType:    model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount:   50.0,
				poolID:      "pool-456",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid task type",
			args: args{
				name:        "Invalid Task Type",
				description: "This task has an invalid type.",
				taskType:    model.TaskType_TASK_TYPE_UNSPECIFIED,
				minAmount:   0.0,
				poolID:      "",
			},
			want: &Task{
				Task: model.Task{
					Name:        "Invalid Task Type",
					Description: "This task has an invalid type.",
					Type:        model.TaskType_TASK_TYPE_UNSPECIFIED,
					Criteria: &model.TaskCriteria{
						MinTransactionAmount: 0.0,
						PoolId:               "",
					},
					Status: model.TaskStatus_TASK_STATUS_ACTIVE,
				},
			},
			wantErr: false,
		},
		{
			name: "Zero minimum amount",
			args: args{
				name:        "Zero Min Amount",
				description: "This task has zero minimum amount.",
				taskType:    model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount:   0.0,
				poolID:      "pool-789",
			},
			want: &Task{
				Task: model.Task{
					Name:        "Zero Min Amount",
					Description: "This task has zero minimum amount.",
					Type:        model.TaskType_TASK_TYPE_ONBOARDING,
					Criteria: &model.TaskCriteria{
						MinTransactionAmount: 0.0,
						PoolId:               "pool-789",
					},
					Status: model.TaskStatus_TASK_STATUS_ACTIVE,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.name, tt.args.description, tt.args.taskType, tt.args.minAmount, tt.args.poolID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_CalculateProgress(t1 *testing.T) {
	type fields struct {
		Name      string
		taskType  model.TaskType
		minAmount float64
		poolID    string
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Task not completed - insufficient amount",
			fields: fields{
				Name:      "Test Task",
				taskType:  model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount: 100.0,
				poolID:    "pool-123",
			},
			args: args{
				amount: 50.0,
			},
			want: 50, // Progress is partial
		},
		{
			name: "Task completed - exact amount",
			fields: fields{
				Name:      "Test Task",
				taskType:  model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount: 100.0,
				poolID:    "pool-123",
			},
			args: args{
				amount: 100.0,
			},
			want: 100, // Progress is 100% complete
		},
		{
			name: "Task completed - excess amount",
			fields: fields{
				Name:      "Test Task",
				taskType:  model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount: 100.0,
				poolID:    "pool-123",
			},
			args: args{
				amount: 150.0,
			},
			want: 100, // Progress is capped at 100%
		},
		{
			name: "Task not started - zero amount",
			fields: fields{
				Name:      "Test Task",
				taskType:  model.TaskType_TASK_TYPE_ONBOARDING,
				minAmount: 100.0,
				poolID:    "pool-123",
			},
			args: args{
				amount: 0.0,
			},
			want: 0, // No progress made
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t, err := NewTask(tt.fields.Name, "", tt.fields.taskType, tt.fields.minAmount, tt.fields.poolID)
			if err != nil {
				t1.Errorf("NewTask() error = %v", err)
				return
			}
			if got := t.CalculateProgress(tt.args.amount); got != tt.want {
				t1.Errorf("CalculateProgress() = %v, want %v", got, tt.want)
			}
		})
	}
}
