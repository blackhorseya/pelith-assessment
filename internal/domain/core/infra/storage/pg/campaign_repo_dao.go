package pg

import (
	"encoding/json"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CampaignDAO 定義 campaigns 表對應的結構
type CampaignDAO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
	Mode        int32     `db:"mode"`
	Status      int32     `db:"status"`
}

// ToBizModel 將 DAO 轉換為 biz.Campaign
func (dao *CampaignDAO) ToBizModel(tasks []*biz.Task) *biz.Campaign {
	return &biz.Campaign{
		Campaign: model.Campaign{
			Id:          dao.ID,
			Name:        dao.Name,
			Description: dao.Description,
			StartTime:   timestamppb.New(dao.StartTime),
			EndTime:     timestamppb.New(dao.EndTime),
			Mode:        model.CampaignMode(dao.Mode),
			Status:      model.CampaignStatus(dao.Status),
		},
		Tasks: tasks,
	}
}

// FromBizModelToDAO 將 biz.Campaign 轉換為 DAO
func FromBizModelToDAO(campaign *biz.Campaign) *CampaignDAO {
	return &CampaignDAO{
		ID:          campaign.Id,
		Name:        campaign.Name,
		Description: campaign.Description,
		StartTime:   campaign.StartTime.AsTime(),
		EndTime:     campaign.EndTime.AsTime(),
		Mode:        int32(campaign.Mode),
		Status:      int32(campaign.Status),
	}
}

// TaskDAO 定義 tasks 表對應的結構
type TaskDAO struct {
	ID          string `db:"id"`
	CampaignID  string `db:"campaign_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Type        int32  `db:"type"`
	Criteria    string `db:"criteria"` // JSON 格式存儲的 criteria
	Status      int32  `db:"status"`
}

// ToBizModel 將 DAO 轉換為 biz.Task
func (dao *TaskDAO) ToBizModel() (*biz.Task, error) {
	var criteria model.TaskCriteria
	if err := json.Unmarshal([]byte(dao.Criteria), &criteria); err != nil {
		return nil, err
	}

	return &biz.Task{
		Task: model.Task{
			Id:          dao.ID,
			Name:        dao.Name,
			Description: dao.Description,
			Type:        model.TaskType(dao.Type),
			Criteria:    &criteria,
			Status:      model.TaskStatus(dao.Status),
		},
	}, nil
}

// FromBizTaskToDAO 將 biz.Task 轉換為 DAO
func FromBizTaskToDAO(task *biz.Task) (*TaskDAO, error) {
	criteriaBytes, err := json.Marshal(task.Criteria)
	if err != nil {
		return nil, err
	}

	return &TaskDAO{
		ID:          task.Id,
		CampaignID:  task.CampaignID,
		Name:        task.Name,
		Description: task.Description,
		Type:        int32(task.Type),
		Criteria:    string(criteriaBytes),
		Status:      int32(task.Status),
	}, nil
}
