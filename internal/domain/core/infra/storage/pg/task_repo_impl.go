package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
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
