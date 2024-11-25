package pg

import (
	"errors"
	"fmt"

	"github.com/blackhorseya/pelith-assessment/scripts/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

func migrateUp(db *sqlx.DB, name string) error {
	// Initialize the PostgreSQL driver
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{
		MigrationsTable:       name + "_migrations",
		MultiStatementEnabled: true,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize postgres driver: %w", err)
	}

	// Initialize the iofs driver with the embedded migrations
	sourceDriver, err := iofs.New(migrations.MigrationsFS, name)
	if err != nil {
		return fmt.Errorf("failed to initialize iofs driver: %w", err)
	}

	// Create a new migrate instance
	migration, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Run the migrations
	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
