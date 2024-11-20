//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type (
	// TaskCreator is used to create a new task.
	TaskCreator interface {
		// Create is used to create a new task.
		Create(c context.Context, task *model.Task) error
	}
)
