//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// ListUserCondition specifies filters for querying users.
type ListUserCondition struct {
	Keyword *string // Optional search keyword for user names or emails
	Limit   int     // Maximum number of users to retrieve
	Offset  int     // Offset for pagination
}

// UserGetter defines the interface for retrieving user data.
type UserGetter interface {
	// GetByID retrieves a user by their ID.
	GetByID(c context.Context, id string) (*model.User, error)

	// List retrieves users based on specified conditions.
	List(c context.Context, cond ListUserCondition) (item []*model.User, total int, err error)
}

// UserQueryStore defines the interface for storing user query data.
type UserQueryStore struct {
}

// NewUserQueryStore creates a new UserQueryStore instance.
func NewUserQueryStore() *UserQueryStore {
	return &UserQueryStore{}
}

// GetTasksStatus retrieves the status of tasks for a user.
func (s *UserQueryStore) GetTasksStatus(c context.Context, address string) (interface{}, error) {
	// TODO: 2024/11/26|sean|implement GetTasksStatus
	return nil, errors.New("not implemented")
}
