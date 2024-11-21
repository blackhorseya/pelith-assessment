//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
)

type (
	// TaskCreator is used to create a new task.
	TaskCreator interface {
		// Create is used to create a new task.
		Create(c context.Context, task *biz.Task, campaignID string) error
	}

	TaskUpdater interface {
		// Update is used to update a task.
		Update(c context.Context, task *biz.Task) error
	}
)

// AddTaskHandler 用於處理 Task 相關的 Command
type AddTaskHandler struct {
	campaignService biz.CampaignService
	campaignGetter  query.CampaignGetter
	taskService     biz.TaskService
	taskCreator     TaskCreator
}

// NewAddTaskHandler 用於建立 AddTaskHandler
func NewAddTaskHandler(
	service biz.CampaignService,
	taskService biz.TaskService,
	taskCreator TaskCreator,
) *AddTaskHandler {
	return &AddTaskHandler{
		campaignService: service,
		taskService:     taskService,
		taskCreator:     taskCreator,
	}
}

func (h *AddTaskHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	cmd, ok := msg.(AddTaskCommand)
	if !ok {
		return "", errors.New("invalid command type for AddTaskHandler")
	}

	err := cmd.Validate()
	if err != nil {
		return "", err
	}

	// TODO: 2024/11/21|sean|implement me
	return "", errors.New("implement me")
}
