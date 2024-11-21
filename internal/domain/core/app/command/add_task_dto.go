package command

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
