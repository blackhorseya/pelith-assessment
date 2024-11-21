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
	// 開啟事務
	tx, err := i.rw.BeginTxx(c, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // 回滾事務
			panic(p)          // 重拋 panic
		} else if err != nil {
			_ = tx.Rollback() // 回滾事務
		} else {
			err = tx.Commit() // 提交事務
		}
	}()

	// 插入 Campaign 並返回生成的 ID
	campaignQuery := `
		INSERT INTO campaigns (name, description, start_time, end_time, mode, status, created_at, updated_at)
		VALUES (:name, :description, :start_time, :end_time, :mode, :status, NOW(), NOW())
		RETURNING id
	`

	campaignParams := map[string]interface{}{
		"name":        campaign.Name,
		"description": campaign.Description,
		"start_time":  campaign.StartTime,
		"end_time":    campaign.EndTime,
		"mode":        campaign.Mode,
		"status":      campaign.Status,
	}

	var campaignID string
	err = tx.QueryRowxContext(c, campaignQuery, campaignParams).Scan(&campaignID)
	if err != nil {
		return err
	}

	// 更新 Campaign 的 ID
	campaign.Id = campaignID

	// 插入 Tasks 並返回生成的 ID
	taskQuery := `
		INSERT INTO tasks (campaign_id, name, description, type, criteria, status, created_at, updated_at)
		VALUES (:campaign_id, :name, :description, :type, :criteria, :status, NOW(), NOW())
		RETURNING id
	`

	for _, task := range campaign.Tasks {
		taskParams := map[string]interface{}{
			"campaign_id": campaignID,
			"name":        task.Name,
			"description": task.Description,
			"type":        task.Type,
			"criteria":    task.Criteria, // 假設是 JSONB 對象
			"status":      task.Status,
		}

		var taskID string
		err = tx.QueryRowxContext(c, taskQuery, taskParams).Scan(&taskID)
		if err != nil {
			return err
		}

		// 更新 Task 的 ID
		task.Id = taskID
	}

	return nil
}
