package pg

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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

//nolint:funlen // it's okay
func (i *campaignRepoImpl) Create(c context.Context, campaign *biz.Campaign) error {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	// 開啟事務
	tx, err := i.rw.BeginTxx(timeout, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 插入 Campaign 並返回生成的 ID
	campaignQuery := `
		INSERT INTO campaigns (name, description, start_time, end_time, mode, status, created_at, updated_at)
		VALUES (:name, :description, :start_time, :end_time, :mode, :status, NOW(), NOW())
		RETURNING id
	`

	type CampaignParams struct {
		Name        string    `db:"name"`
		Description string    `db:"description"`
		StartTime   time.Time `db:"start_time"`
		EndTime     time.Time `db:"end_time"`
		Mode        int       `db:"mode"`
		Status      int       `db:"status"`
	}

	campaignParams := CampaignParams{
		Name:        campaign.Name,
		Description: campaign.Description,
		StartTime:   campaign.StartTime.AsTime(),
		EndTime:     campaign.EndTime.AsTime(),
		Mode:        int(campaign.Mode),
		Status:      int(campaign.Status),
	}

	var campaignID string
	stmt, err := tx.PrepareNamedContext(timeout, campaignQuery)
	if err != nil {
		ctx.Error("failed to prepare named statement", zap.Error(err))
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowxContext(timeout, campaignParams).Scan(&campaignID)
	if err != nil {
		ctx.Error("failed to insert campaign", zap.Error(err))
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

	type TaskParams struct {
		CampaignID  string `db:"campaign_id"`
		Name        string `db:"name"`
		Description string `db:"description"`
		Type        int    `db:"type"`
		Criteria    string `db:"criteria"`
		Status      int    `db:"status"`
	}

	taskStmt, err := tx.PrepareNamedContext(timeout, taskQuery)
	if err != nil {
		ctx.Error("failed to prepare named statement for tasks", zap.Error(err))
		return err
	}
	defer taskStmt.Close()

	for _, task := range campaign.Tasks {
		criteria, err2 := json.Marshal(task.Criteria)
		if err2 != nil {
			ctx.Error("failed to marshal task criteria", zap.Error(err2))
			return err2
		}

		taskParams := TaskParams{
			CampaignID:  campaignID,
			Name:        task.Name,
			Description: task.Description,
			Type:        int(task.Type),
			Criteria:    string(criteria),
			Status:      int(task.Status),
		}

		var taskID string
		err = taskStmt.QueryRowxContext(timeout, taskParams).Scan(&taskID)
		if err != nil {
			ctx.Error("failed to insert task", zap.Error(err))
			return err
		}

		// 更新 Task 的 ID
		task.Id = taskID
	}

	return nil
}

func (i *campaignRepoImpl) GetByID(c context.Context, id string) (*model.Campaign, error) {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	// Query to fetch the campaign by ID
	campaignQuery := `
		SELECT id, name, description, start_time, end_time, mode, status
		FROM campaigns
		WHERE id = $1
	`

	var campaign model.Campaign

	err := i.rw.GetContext(timeout, &campaign, campaignQuery, id)
	if err != nil {
		ctx.Error("failed to fetch campaign by id", zap.Error(err))
		return nil, err
	}

	// Query to fetch tasks associated with the campaign
	taskQuery := `
		SELECT id, name, description, type, criteria, status
		FROM tasks
		WHERE campaign_id = $1
	`

	var tasks []*model.Task

	err = i.rw.SelectContext(timeout, &tasks, taskQuery, id)
	if err != nil {
		ctx.Error("failed to fetch tasks for campaign", zap.Error(err))
		return nil, err
	}

	campaign.Tasks = make([]string, len(tasks))
	for idx, task := range tasks {
		campaign.Tasks[idx] = task.Id
	}

	return &campaign, nil
}

func (i *campaignRepoImpl) List(
	c context.Context,
	cond query.ListCampaignCondition,
) (items []*model.Campaign, total int, err error) {
	// TODO: 2024/11/21|sean|implement me
	panic("implement me")
}
