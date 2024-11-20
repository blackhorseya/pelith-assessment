//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"

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
