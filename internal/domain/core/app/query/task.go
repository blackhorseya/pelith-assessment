//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

// ListTaskCondition is the condition to list the task.
type ListTaskCondition struct {
}

// TaskGetter is used to get the task.
type TaskGetter interface {
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
func (s *TaskQueryService) GetTaskStatus(c context.Context, address string) ([]*biz.Task, error) {
	// fetch tasks by address

	// for loop tasks and calculate progress

	// TODO: 2024/11/22|sean|implement the logic
	return nil, errors.New("not implemented")
}
