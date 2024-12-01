package command

import (
	"errors"
	"strconv"
)

// TaskCommand 用於描述 Task 的輸入參數
type TaskCommand struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        int32   `json:"type"`
	MinAmount   float64 `json:"min_amount"`
	PoolID      string  `json:"pool_id"`
}

// AddTaskCommand 用於新增 Tasks 至指定 Campaign
type AddTaskCommand struct {
	CampaignID string        `json:"campaign_id"`
	Tasks      []TaskCommand `json:"tasks"`
}

func (cmd AddTaskCommand) Key() int {
	return addTaskCommandKey
}

func (cmd AddTaskCommand) Validate() error {
	// 檢查 CampaignID 是否有效
	if cmd.CampaignID == "" {
		return errors.New("campaign ID is required")
	}

	// 檢查 Tasks 是否存在
	if len(cmd.Tasks) == 0 {
		return errors.New("at least one task is required")
	}

	// 檢查每個 TaskCommand 的屬性
	for i, task := range cmd.Tasks {
		if task.Name == "" {
			return errors.New("task name is required for task " + strconv.Itoa(i))
		}
		if task.Type <= 0 {
			return errors.New("task type must be valid for task " + strconv.Itoa(i))
		}
		if task.MinAmount < 0 {
			return errors.New("minimum transaction amount cannot be negative for task " + strconv.Itoa(i))
		}
	}

	return nil
}
