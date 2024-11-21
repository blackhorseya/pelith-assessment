package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type taskServiceImpl struct {
}

// NewTaskService creates a new TaskService instance.
func NewTaskService() biz.TaskService {
	return &taskServiceImpl{}
}

func (i *taskServiceImpl) CreateTask(c context.Context, campaignID string, task *model.Task) (*biz.Task, error) {
	return biz.NewTask(campaignID, task.Name, task.Description, task.Type, task.Criteria)
}
