//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type (
	// UserCreator defines the interface for creating a new user.
	UserCreator interface {
		// Create persists a new user.
		Create(c context.Context, user *model.User) error
	}

	// UserUpdater defines the interface for updating user information.
	UserUpdater interface {
		// Update modifies an existing user's data.
		Update(c context.Context, user *model.User) error

		// IncrementPoints adds points to a user's total.
		IncrementPoints(c context.Context, userID string, points int64) error
	}
)
