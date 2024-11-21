//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
)

type (
	// TaskCreator is used to create a new task.
	TaskCreator interface {
		// Create is used to create a new task.
		Create(c context.Context, task *biz.Task, campaignID string) error
	}
)

// AddTaskHandler 用於處理 Task 相關的 Command
type AddTaskHandler struct {
	service     biz.CampaignService
	taskCreator TaskCreator
}

// NewAddTaskHandler 用於建立 AddTaskHandler
func NewAddTaskHandler(service biz.CampaignService, taskCreator TaskCreator) *AddTaskHandler {
	return &AddTaskHandler{service: service, taskCreator: taskCreator}
}

func (h *AddTaskHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	// TODO: 2024/11/21|sean|處理新增 Task 的邏輯
	panic("implement me")
}
