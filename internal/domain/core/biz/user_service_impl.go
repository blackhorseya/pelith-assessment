package biz

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

type userServiceImpl struct {
}

// NewUserService creates a new UserService instance.
func NewUserService() biz.UserService {
	return &userServiceImpl{}
}

func (i *userServiceImpl) GetUserTaskListByAddress(c context.Context, address string) (*biz.User, error) {
	// TODO: 2024/11/26|sean|implement GetUserTaskListByAddress
	panic("implement me")
}
