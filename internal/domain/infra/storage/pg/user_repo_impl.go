package pg

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // import migration files
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import postgres driver
)

// UserRepoImpl is a postgres implementation of UserRepo
type UserRepoImpl struct {
}

// NewUserRepo creates a new UserRepoImpl
func NewUserRepo(rw *sqlx.DB) (*UserRepoImpl, error) {
	driver, err := postgres.WithInstance(rw.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance(migrationFolder+"/user", "postgres", driver)
	if err != nil {
		return nil, err
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &UserRepoImpl{}, nil
}
