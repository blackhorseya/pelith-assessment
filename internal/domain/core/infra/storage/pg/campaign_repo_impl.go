package pg

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

type campaignRepoImpl struct {
	rw *sqlx.DB
}

// NewCampaignRepo is used to create a new campaignRepoImpl.
func NewCampaignRepo(rw *sqlx.DB) (command.CampaignCreator, error) {
	driver, err := postgres.WithInstance(rw.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance(migrationFolder+"/campaign", "postgres", driver)
	if err != nil {
		return nil, err
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &campaignRepoImpl{
		rw: rw,
	}, nil
}

func (i *campaignRepoImpl) Create(c context.Context, campaign *biz.Campaign) error {
	// TODO: 2024/11/20|sean|implement create campaign
	panic("implement me")
}
