package biz

import (
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// Task is an aggregate root that represents the task.
type Task struct {
	model.Task

	CampaignID string
	Progress   int
}

// NewTask creates a new Task aggregate.
func NewTask(
	name, description string,
	taskType model.TaskType,
	criteria *model.TaskCriteria,
) (*Task, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if criteria == nil {
		return nil, errors.New("criteria is required")
	}

	return &Task{
		Task: model.Task{
			Id:          "", // id will be generated by the repository
			Name:        name,
			Description: description,
			Type:        taskType,
			Criteria:    criteria,
			Status:      model.TaskStatus_TASK_STATUS_ACTIVE,
		},
	}, nil
}

// Evaluate checks whether a task is completed based on the given inputs.
func (t *Task) Evaluate(transactionAmount float64, poolID string) (bool, error) {
	if t.Status != model.TaskStatus_TASK_STATUS_ACTIVE {
		return false, errors.New("task is not active")
	}

	if transactionAmount < t.Criteria.MinTransactionAmount {
		return false, nil
	}

	if t.Criteria != nil && t.Criteria.PoolId != "" && poolID != t.Criteria.PoolId {
		return false, nil
	}

	return true, nil
}

// Deactivate marks the task as inactive.
func (t *Task) Deactivate() {
	t.Status = model.TaskStatus_TASK_STATUS_INACTIVE
}
