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

// TaskCommand 用於描述 Task 的輸入參數
type TaskCommand struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        int     `json:"type"`
	MinAmount   float64 `json:"min_amount"`
	PoolID      string  `json:"pool_id"`
}

// AddTaskCommand 用於新增 Tasks 至指定 Campaign
type AddTaskCommand struct {
	CampaignID string        `json:"campaign_id"`
	Tasks      []TaskCommand `json:"tasks"`
}

// TaskCommandHandler 用於處理 Task 相關的 Command
type TaskCommandHandler struct {
	service     biz.CampaignService
	taskCreator TaskCreator
}

func (h *TaskCommandHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	// TODO: 2024/11/21|sean|處理新增 Task 的邏輯
	panic("implement me")
}
