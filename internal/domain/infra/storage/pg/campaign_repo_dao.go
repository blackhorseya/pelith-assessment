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
	PoolID      string    `db:"pool_id"`
}

// ToBizModel 將 DAO 轉換為 biz.Campaign
func (dao *CampaignDAO) ToBizModel(tasks []*biz.Task) *biz.Campaign {
	campaign, _ := biz.NewCampaign(dao.Name, dao.StartTime, dao.PoolID)
	campaign.Id = dao.ID
	campaign.Description = dao.Description
	campaign.EndTime = timestamppb.New(dao.EndTime)
	campaign.Mode = model.CampaignMode(dao.Mode)
	campaign.Status = model.CampaignStatus(dao.Status)
	for _, task := range tasks {
		_ = campaign.AddTask(task)
	}

	return campaign
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
		PoolID:      campaign.PoolId,
	}
}

// TaskDAO 定義 tasks 表對應的結構
type TaskDAO struct {
	ID          string    `db:"id"`
	CampaignID  string    `db:"campaign_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Type        int32     `db:"type"`
	Criteria    string    `db:"criteria"` // JSON 格式存儲的 criteria
	Status      int32     `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
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
		CampaignID: dao.CampaignID,
		Progress:   0,
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

// RewardDAO 定義 rewards 表對應的結構
type RewardDAO struct {
	ID          string     `db:"id"`
	UserAddress string     `db:"user_address"`
	CampaignID  string     `db:"campaign_id"`
	Points      int64      `db:"points"`
	RedeemedAt  *time.Time `db:"redeemed_at"` // 可為 NULL
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}

// FromModelRewardToDAO 將 biz.Reward 轉換為 DAO
func FromModelRewardToDAO(reward *model.Reward) *RewardDAO {
	redeemedAt := new(time.Time)
	if !redeemedAt.IsZero() {
		*redeemedAt = reward.RedeemedAt.AsTime()
	}

	return &RewardDAO{
		ID:          reward.Id,
		UserAddress: reward.UserAddress,
		CampaignID:  reward.CampaignId,
		Points:      reward.Points,
		RedeemedAt:  redeemedAt,
	}
}

// ToModel 將 RewardDAO 轉換為 model.Reward
func (dao *RewardDAO) ToModel() *model.Reward {
	return &model.Reward{
		Id:          dao.ID,
		UserAddress: dao.UserAddress,
		CampaignId:  dao.CampaignID,
		Points:      dao.Points,
		RedeemedAt:  timestamppb.New(*dao.RedeemedAt),
	}
}

// ToAggregate 將 RewardDAO 轉換為 biz.Reward
func (dao *RewardDAO) ToAggregate() *biz.Reward {
	return &biz.Reward{
		Reward: model.Reward{
			Id:          dao.ID,
			UserAddress: dao.UserAddress,
			CampaignId:  dao.CampaignID,
			Points:      dao.Points,
			RedeemedAt:  timestamppb.New(*dao.RedeemedAt),
		},
	}
}
