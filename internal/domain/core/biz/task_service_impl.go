package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

type taskServiceImpl struct {
}

// NewTaskService creates a new TaskService instance.
func NewTaskService() biz.TaskService {
	return &taskServiceImpl{}
}

func (t taskServiceImpl) CreateTask(
	c context.Context,
	campaign *biz.Campaign,
	name, description string,
	taskType model.TaskType,
	minAmount float64,
	poolID string,
) (*biz.Task, error) {
	ctx := contextx.WithContext(c)

	task, err := biz.NewTask(name, description, taskType, &model.TaskCriteria{
		MinTransactionAmount: minAmount,
		PoolId:               poolID,
	})
	if err != nil {
		ctx.Error("failed to create task", zap.Error(err))
		return nil, err
	}

	err = campaign.AddTask(task)
	if err != nil {
		ctx.Error("failed to add task to campaign", zap.Error(err))
		return nil, err
	}

	return task, nil
}
