package biz

import (
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// Task is an aggregate root that represents the task.
type Task struct {
	model.Task
}

// NewTask creates a new Task aggregate.
func NewTask(id, name, description string, taskType model.TaskType, criteria *model.TaskCriteria) (*Task, error) {
	if id == "" || name == "" {
		return nil, errors.New("task ID and name are required")
	}

	return &Task{
		Task: model.Task{
			Id:          id,
			Name:        name,
			Description: description,
			Type:        taskType,
			Criteria:    criteria,
			Status:      model.TaskStatus_TASK_STATUS_ACTIVE,
		},
	}, nil
}
