//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

// ListTaskCondition is the condition to list the task.
type ListTaskCondition struct {
}

// TaskGetter is used to get the task.
type TaskGetter interface {
	ListTask(c context.Context, cond ListTaskCondition) (items []*biz.Task, total int, err error)
}

// TaskQueryService is the service for task query.
type TaskQueryService struct {
	taskGetter TaskGetter
}

// NewTaskQueryService is used to create a new TaskQueryService.
func NewTaskQueryService(taskGetter TaskGetter) *TaskQueryService {
	return &TaskQueryService{taskGetter: taskGetter}
}

// GetTaskStatus is used to get the task status.
func (s *TaskQueryService) GetTaskStatus(
	c context.Context,
	address string,
	campaignID string,
) ([]*biz.Task, error) {
	// TODO: 2024/11/22|sean|fetch tasks by address
	// fetch tasks by address
	var tasks []*biz.Task

	// for loop tasks and calculate progress
	for _, task := range tasks {
		// TODO: 2024/11/22|sean|pass the correct amount
		task.Progress = task.CalculateProgress(0)
	}

	return tasks, nil
}
