package biz

import (
	"testing"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

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
