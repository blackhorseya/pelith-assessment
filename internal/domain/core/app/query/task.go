//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// ListTaskCondition is the condition to list the task.
type ListTaskCondition struct {
	CampaignID string
	Status     model.TaskStatus
}

// TaskGetter is used to get the task.
type TaskGetter interface {
	ListTask(c context.Context, cond ListTaskCondition) (items []*biz.Task, total int, err error)
}

// TaskQueryService is the service for task query.
type TaskQueryService struct {
	txQueryService *TransactionQueryService
	taskGetter     TaskGetter
}

// NewTaskQueryService is used to create a new TaskQueryService.
func NewTaskQueryService(taskGetter TaskGetter, txQueryService *TransactionQueryService) *TaskQueryService {
	return &TaskQueryService{
		taskGetter:     taskGetter,
		txQueryService: txQueryService,
	}
}

// GetTaskStatus is used to get the task status.
func (s *TaskQueryService) GetTaskStatus(
	c context.Context,
	address string,
	campaignID string,
) ([]*biz.Task, error) {
	tasks, _, err := s.taskGetter.ListTask(c, ListTaskCondition{
		CampaignID: campaignID,
	})
	if err != nil {
		return nil, err
	}

	// for loop tasks and calculate progress
	for _, task := range tasks {
		txAmount, err2 := s.txQueryService.GetTotalSwapAmount(c, address, task.CampaignID)
		if err2 != nil {
			return nil, err2
		}

		task.Progress = task.CalculateProgress(txAmount)
	}

	return tasks, nil
}
