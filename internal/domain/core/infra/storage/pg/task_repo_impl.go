package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// TaskRepoImpl is the implementation of TaskRepo
type TaskRepoImpl struct {
	rw *sqlx.DB
}

// NewTaskRepo creates a new TaskRepoImpl
func NewTaskRepo(rw *sqlx.DB) *TaskRepoImpl {
	return &TaskRepoImpl{rw: rw}
}

// NewTaskCreator creates a new TaskCreator
func NewTaskCreator(impl *TaskRepoImpl) command.TaskCreator {
	return impl
}

// NewTaskGetter creates a new TaskGetter
func NewTaskGetter(impl *TaskRepoImpl) query.TaskGetter {
	return impl
}

func (i *TaskRepoImpl) Create(c context.Context, task *biz.Task) error {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 開啟事務
	tx, err := i.rw.BeginTxx(timeout, nil)
	if err != nil {
		ctx.Error("failed to begin transaction", zap.Error(err))
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			ctx.Error("transaction rolled back due to panic", zap.Any("panic", p))
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
			ctx.Error("transaction rolled back due to error", zap.Error(err))
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				ctx.Error("failed to commit transaction", zap.Error(commitErr))
				err = commitErr
			}
		}
	}()

	// 將 Task 轉換為 TaskDAO
	taskDAO, err := FromBizTaskToDAO(task)
	if err != nil {
		ctx.Error("failed to convert task to DAO", zap.Error(err))
		return err
	}

	// 插入 Task 並返回生成的 ID
	taskQuery := `
		INSERT INTO tasks (campaign_id, name, description, type, criteria, status, created_at, updated_at)
		VALUES (:campaign_id, :name, :description, :type, :criteria, :status, NOW(), NOW())
		RETURNING id
	`
	stmt, err := tx.PrepareNamedContext(timeout, taskQuery)
	if err != nil {
		ctx.Error("failed to prepare named statement", zap.Error(err))
		return err
	}
	defer stmt.Close()

	var taskID string
	err = stmt.QueryRowxContext(timeout, taskDAO).Scan(&taskID)
	if err != nil {
		ctx.Error("failed to insert task", zap.Error(err))
		return err
	}

	// 更新 Task 的 ID
	task.Id = taskID
	ctx.Info("task created successfully", zap.String("task_id", taskID))

	return nil
}

func (i *TaskRepoImpl) ListTask(
	c context.Context,
	cond query.ListTaskCondition,
) (items []*biz.Task, total int, err error) {
	ctx := contextx.WithContext(c)

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 構建查詢語句
	queryBuilder := `
		SELECT id, campaign_id, name, description, type, criteria, status, created_at, updated_at
		FROM tasks
		WHERE 1=1
	`
	countQueryBuilder := `
		SELECT COUNT(*)
		FROM tasks
		WHERE 1=1
	`
	// 用具名結構替代 map
	type queryArgs struct {
		CampaignID string `db:"campaign_id"`
		Status     int32  `db:"status"`
	}
	args := queryArgs{}

	// 根據條件動態構建 SQL 語句和參數
	if cond.CampaignID != "" {
		queryBuilder += " AND campaign_id = :campaign_id"
		countQueryBuilder += " AND campaign_id = :campaign_id"
		args.CampaignID = cond.CampaignID
	}
	if cond.Status != 0 {
		queryBuilder += " AND status = :status"
		countQueryBuilder += " AND status = :status"
		args.Status = int32(cond.Status)
	}

	// 查詢總數
	var totalCount int
	nstmt, err := i.rw.PrepareNamed(countQueryBuilder)
	if err != nil {
		ctx.Error("failed to prepare count query", zap.Error(err))
		return nil, 0, err
	}
	defer nstmt.Close()

	err = nstmt.GetContext(timeout, &totalCount, args)
	if err != nil {
		ctx.Error("failed to fetch total count of tasks", zap.Error(err))
		return nil, 0, err
	}
	total = totalCount

	// 查詢任務列表
	var taskDAOs []TaskDAO
	nstmt, err = i.rw.PrepareNamed(queryBuilder)
	if err != nil {
		ctx.Error("failed to prepare task query", zap.Error(err))
		return nil, 0, err
	}
	defer nstmt.Close()

	err = nstmt.SelectContext(timeout, &taskDAOs, args)
	if err != nil {
		ctx.Error("failed to fetch tasks", zap.Error(err))
		return nil, 0, err
	}

	// 將 TaskDAO 轉換為 biz.Task
	for _, taskDAO := range taskDAOs {
		task, convErr := taskDAO.ToBizModel()
		if convErr != nil {
			ctx.Error("failed to convert task DAO to biz model", zap.Error(convErr))
			return nil, 0, convErr
		}
		items = append(items, task)
	}

	return items, total, nil
}
