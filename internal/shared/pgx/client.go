package pgx

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultConns       = 100
	defaultMaxLifetime = 15 * time.Minute
	defaultTimeout     = 5 * time.Second
	defaultLimit       = 10
	defaultMaxLimit    = 100
)

// NewClient init mysql client.
func NewClient(app *configx.Application) (*sqlx.DB, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelFunc()

	db, err := sqlx.ConnectContext(timeout, "postgres", app.Storage.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(defaultConns)
	db.SetMaxIdleConns(defaultConns)
	db.SetConnMaxLifetime(defaultMaxLifetime)

	return db, nil
}
