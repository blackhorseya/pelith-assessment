package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type userServiceImpl struct {
}

// NewUserService creates a new UserService instance.
func NewUserService() biz.UserService {
	return &userServiceImpl{}
}

func (i *userServiceImpl) RegisterUser(c context.Context, name string, address string) (*model.User, error) {
	// TODO: 2024/11/20|sean|implement me
	panic("implement me")
}

func (i *userServiceImpl) UpdateUserProgress(c context.Context, userID string, taskID string, completed bool) error {
	// TODO: 2024/11/20|sean|implement me
	panic("implement me")
}