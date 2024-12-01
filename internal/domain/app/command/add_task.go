package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
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
	campaignService biz.CampaignService,
	campaignGetter query.CampaignGetter,
	taskService biz.TaskService,
	taskCreator TaskCreator,
) *AddTaskHandler {
	return &AddTaskHandler{
		campaignService: campaignService,
		campaignGetter:  campaignGetter,
		taskService:     taskService,
		taskCreator:     taskCreator,
	}
}

func (h *AddTaskHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	ctx := contextx.WithContext(c)

	// 驗證輸入的命令
	cmd, ok := msg.(AddTaskCommand)
	if !ok {
		return "", errors.New("invalid command type for AddTaskHandler")
	}

	err := cmd.Validate()
	if err != nil {
		ctx.Error("validation failed", zap.Error(err), zap.Any("command", &cmd))
		return "", err
	}

	// 根據 CampaignID 獲取對應的 Campaign
	campaign, err := h.campaignGetter.GetByID(c, cmd.CampaignID)
	if err != nil {
		ctx.Error("failed to fetch campaign", zap.Error(err))
		return "", err
	}

	// 逐一處理 Tasks
	for _, taskCmd := range cmd.Tasks {
		// 使用 TaskService 創建新 Task 並加入 Campaign
		task, err2 := h.taskService.CreateTask(
			c,
			campaign,
			taskCmd.Name,
			taskCmd.Description,
			model.TaskType(taskCmd.Type),
			taskCmd.MinAmount,
			taskCmd.PoolID,
		)
		if err2 != nil {
			ctx.Error("failed to create task", zap.Error(err2))
			return "", err2
		}

		// 保存 Task 到資料庫
		err2 = h.taskCreator.Create(c, task)
		if err2 != nil {
			ctx.Error("failed to save task", zap.Error(err2))
			return "", err2
		}
	}

	// 返回成功訊息
	return campaign.Id, nil
}
