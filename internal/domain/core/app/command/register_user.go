package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
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
