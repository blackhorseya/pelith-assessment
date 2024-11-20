//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
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

// RegisterUserHandler handles the registration of a new user.
type RegisterUserHandler struct {
	biz  biz.UserService
	repo UserCreator
}

// NewRegisterUserHandler creates a new RegisterUserHandler instance.
func NewRegisterUserHandler(b biz.UserService, r UserCreator) RegisterUserHandler {
	return RegisterUserHandler{biz: b, repo: r}
}

// Handle creates a new user in the system.
func (h *RegisterUserHandler) Handle(c context.Context, name string, address string) error {
	// TODO: 2024/11/20|sean|implement me
	panic("implement me")
}
