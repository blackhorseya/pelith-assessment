package pg

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/jmoiron/sqlx"
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

func (i *TaskRepoImpl) Create(c context.Context, task *biz.Task, campaignID string) error {
	// TODO: 2024/11/21|sean|處理新增 Task 的邏輯
	panic("implement me")
}
