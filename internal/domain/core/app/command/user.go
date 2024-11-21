//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
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

// RegisterUserCommand is a command to register a new user.
type RegisterUserCommand struct {
	Username string
	Address  string
}

// UserRegistrationHandler handles the registration of a new user.
type UserRegistrationHandler struct {
	biz  biz.UserService
	repo UserCreator
}

// NewRegisterUserHandler creates a new UserRegistrationHandler instance.
func NewRegisterUserHandler(b biz.UserService, r UserCreator) *UserRegistrationHandler {
	return &UserRegistrationHandler{biz: b, repo: r}
}

func (h *UserRegistrationHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	// TODO: 2024/11/21|sean|Implement the user registration handler.
	panic("implement me")
}
