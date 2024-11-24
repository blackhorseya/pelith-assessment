package pg

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

func migrateUp(db *sqlx.DB, name string) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{
		MigrationsTable:       name + "_migrations",
		MultiStatementEnabled: true,
	})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(migrationFolder+"/"+name, "postgres", driver)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
