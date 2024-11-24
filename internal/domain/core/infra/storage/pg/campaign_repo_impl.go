package pg

import (
	"context"
	"sync"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	instance *CampaignRepoImpl
	mu       sync.Mutex
)

// CampaignRepoImpl is the implementation of CampaignRepo.
type CampaignRepoImpl struct {
	rw *sqlx.DB
}

func NewCampaignRepo(rw *sqlx.DB) (*CampaignRepoImpl, error) {
	err := migrateUp(rw, "campaign")
	if err != nil {
		return nil, err
	}

	return &CampaignRepoImpl{rw: rw}, nil
}

// NewCampaignCreator is used to create a new CampaignCreator.
func NewCampaignCreator(impl *CampaignRepoImpl) (command.CampaignCreator, error) {
	return impl, nil
}

// NewCampaignGetter is used to create a new CampaignGetter.
func NewCampaignGetter(impl *CampaignRepoImpl) (query.CampaignGetter, error) {
	return impl, nil
}

func (i *CampaignRepoImpl) Create(c context.Context, campaign *biz.Campaign) error {
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

	campaignParams := FromBizModelToDAO(campaign)
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

	taskStmt, err := tx.PrepareNamedContext(timeout, taskQuery)
	if err != nil {
		ctx.Error("failed to prepare named statement for tasks", zap.Error(err))
		return err
	}
	defer taskStmt.Close()

	for _, task := range campaign.Tasks {
		task.CampaignID = campaignID
		taskParams, err2 := FromBizTaskToDAO(task)
		if err2 != nil {
			ctx.Error("failed to convert task to DAO", zap.Error(err2))
			return err2
		}

		var taskID string
		err2 = taskStmt.QueryRowxContext(timeout, taskParams).Scan(&taskID)
		if err2 != nil {
			ctx.Error("failed to insert task", zap.Error(err2))
			return err2
		}

		// 更新 Task 的 ID
		task.Id = taskID
	}

	return nil
}

func (i *CampaignRepoImpl) GetByID(c context.Context, id string) (*biz.Campaign, error) {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	// 查詢 Campaign 資料
	var campaignDAO CampaignDAO
	campaignQuery := `
		SELECT id, name, description, start_time, end_time, mode, status
		FROM campaigns
		WHERE id = $1
	`
	err := i.rw.GetContext(timeout, &campaignDAO, campaignQuery, id)
	if err != nil {
		ctx.Error("failed to fetch campaign", zap.Error(err))
		return nil, err
	}

	// 查詢 Tasks 資料
	var taskDAOs []TaskDAO
	taskQuery := `
		SELECT id, campaign_id, name, description, type, criteria, status
		FROM tasks
		WHERE campaign_id = $1
	`
	err = i.rw.SelectContext(timeout, &taskDAOs, taskQuery, id)
	if err != nil {
		ctx.Error("failed to fetch tasks", zap.Error(err))
		return nil, err
	}

	// 將 Tasks 轉換為 biz.Task
	var tasks []*biz.Task
	for _, taskDAO := range taskDAOs {
		task, err2 := taskDAO.ToBizModel()
		if err2 != nil {
			ctx.Error("failed to convert task DAO to biz model", zap.Error(err2))
			return nil, err2
		}
		tasks = append(tasks, task)
	}

	// 將 CampaignDAO 轉換為 biz.Campaign
	return campaignDAO.ToBizModel(tasks), nil
}

func (i *CampaignRepoImpl) List(
	c context.Context,
	cond query.ListCampaignCondition,
) (items []*biz.Campaign, total int, err error) {
	// TODO: 2024/11/21|sean|implement me
	panic("implement me")
}
