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
