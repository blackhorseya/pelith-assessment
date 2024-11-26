package pg

import (
	"context"
	"errors"
	"sync"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
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

// NewCampaignUpdater is used to create a new CampaignUpdater.
func NewCampaignUpdater(impl *CampaignRepoImpl) (command.CampaignUpdater, error) {
	return impl, nil
}

// NewCampaignDeleter is used to create a new CampaignDeleter.
func NewCampaignDeleter(impl *CampaignRepoImpl) (command.CampaignDeleter, error) {
	return impl, nil
}

func (i *CampaignRepoImpl) CleanReward(c context.Context, campaignID string) error {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	stmt := `
		DELETE FROM rewards
		WHERE campaign_id = $1
	`

	_, err := i.rw.ExecContext(timeout, stmt, campaignID)
	if err != nil {
		ctx.Error("failed to clean rewards", zap.Error(err), zap.String("campaign_id", campaignID))
		return err
	}

	return nil
}

func (i *CampaignRepoImpl) DistributeReward(c context.Context, reward *model.Reward) error {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	if reward == nil {
		return errors.New("reward is nil")
	}

	rewardQuery := `
		INSERT INTO rewards (user_address, campaign_id, points, redeemed_at, created_at, updated_at)
		VALUES (:user_address, :campaign_id, :points, :redeemed_at, NOW(), NOW())
		RETURNING id
	`
	stmt, err := i.rw.PrepareNamedContext(timeout, rewardQuery)
	if err != nil {
		ctx.Error("failed to prepare named statement", zap.Error(err))
		return err
	}
	defer stmt.Close()

	params := FromModelRewardToDAO(reward)
	var rewardID string
	err = stmt.QueryRowxContext(timeout, params).Scan(&rewardID)
	if err != nil {
		ctx.Error("failed to insert reward", zap.Error(err), zap.Any("params", &params))
		return err
	}

	reward.Id = rewardID
	return nil
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
		INSERT INTO campaigns (name, description, start_time, end_time, mode, status, pool_id, created_at, updated_at)
		VALUES (:name, :description, :start_time, :end_time, :mode, :status, :pool_id, NOW(), NOW())
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

	for _, task := range campaign.Tasks() {
		taskParams, err2 := FromBizTaskToDAO(task)
		taskParams.CampaignID = campaignID
		if err2 != nil {
			ctx.Error("failed to convert task to DAO", zap.Error(err2))
			return err2
		}

		var taskID string
		err2 = taskStmt.QueryRowxContext(timeout, taskParams).Scan(&taskID)
		if err2 != nil {
			ctx.Error("failed to insert task", zap.Error(err2), zap.Any("params", &taskParams))
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
		SELECT id, name, description, start_time, end_time, mode, status, pool_id
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
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	// Validate limit
	const defaultMaxLimit = 1000
	if cond.Limit <= 0 || cond.Limit > defaultMaxLimit {
		cond.Limit = defaultMaxLimit
	}

	// Query to count total campaigns
	countQuery := `
		SELECT COUNT(*) 
		FROM campaigns
	`
	err = i.rw.GetContext(timeout, &total, countQuery)
	if err != nil {
		ctx.Error("failed to count campaigns", zap.Error(err))
		return nil, 0, err
	}

	if total == 0 {
		// No campaigns to return
		return []*biz.Campaign{}, 0, nil
	}

	// Query to fetch campaigns
	campaignQuery := `
		SELECT id, name, description, start_time, end_time, mode, status, pool_id
		FROM campaigns
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	var campaignDAOs []CampaignDAO
	err = i.rw.SelectContext(timeout, &campaignDAOs, campaignQuery, cond.Limit, cond.Offset)
	if err != nil {
		ctx.Error("failed to fetch campaigns", zap.Error(err))
		return nil, 0, err
	}

	// Fetch tasks for each campaign
	items = make([]*biz.Campaign, 0, len(campaignDAOs))
	for _, campaignDAO := range campaignDAOs {
		// Query to fetch tasks for the campaign
		var taskDAOs []TaskDAO
		taskQuery := `
			SELECT id, campaign_id, name, description, type, criteria, status
			FROM tasks
			WHERE campaign_id = $1
		`
		err = i.rw.SelectContext(timeout, &taskDAOs, taskQuery, campaignDAO.ID)
		if err != nil {
			ctx.Error(
				"failed to fetch tasks for campaign",
				zap.String("campaign_id", campaignDAO.ID),
				zap.Error(err),
			)
			return nil, 0, err
		}

		// Convert tasks to biz models
		var tasks []*biz.Task
		for _, taskDAO := range taskDAOs {
			task, err2 := taskDAO.ToBizModel()
			if err2 != nil {
				ctx.Error("failed to convert task DAO to biz model", zap.Error(err2))
				return nil, 0, err2
			}
			tasks = append(tasks, task)
		}

		// Convert campaign to biz model and add to list
		campaign := campaignDAO.ToBizModel(tasks)
		items = append(items, campaign)
	}

	return items, total, nil
}
