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
	// GetByID is used to get a task by id.
	GetByID(c context.Context, id string) (*biz.Task, error)

	// List is used to list the task.
	List(c context.Context, cond ListTaskCondition) (items []*biz.Task, total int, err error)
}
