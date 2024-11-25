package biz

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/composite"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"go.uber.org/zap"
)

type backtestServiceImpl struct {
	bus    eventx.EventBus
	txRepo *composite.TransactionCompositeRepoImpl
}

// NewBacktestService creates a new BacktestService instance.
func NewBacktestService(txRepo *composite.TransactionCompositeRepoImpl) biz.BacktestService {
	return &backtestServiceImpl{
		txRepo: txRepo,
	}
}

//nolint:gocognit // ignore cognitive complexity
func (i *backtestServiceImpl) RunBacktest(
	c context.Context,
	campaign *biz.Campaign,
	resultCh chan<- *model.Reward,
) error {
	ctx := contextx.WithContext(c)

	// 1. 獲取交易記錄
	txCh := make(chan *model.Transaction)
	var err error
	err = i.txRepo.GetSwapTxByPoolAddress(c, campaign.PoolId, query.ListTransactionCondition{
		StartTime: campaign.StartTime.AsTime(),
		EndTime:   campaign.EndTime.AsTime(),
	}, txCh)

	// 2. 準備累積數據
	userSwapVolume := make(map[string]float64) // 用戶的交易量
	totalSwapVolume := 0.0                     // 總交易量

	// 3. 處理交易記錄
	for _, tx := range transactionList {
		for _, task := range campaign.Tasks() {
			// 處理 Onboarding Task
			if task.Type == model.TaskType_TASK_TYPE_ONBOARDING &&
				float64(tx.GetTransaction().Amount) >= task.Criteria.MinTransactionAmount {
				// 發放 Onboarding Task 獎勵
				reward := &model.Reward{
					Id:         "", // 生成唯一 ID
					UserId:     tx.GetTransaction().FromAddress,
					CampaignId: campaign.Id,
					Points:     100, // 固定獎勵點數
				}

				// 發送到結果通道
				select {
				case resultCh <- reward:
					ctx.Info(
						"Onboarding Task reward sent",
						zap.String("user", tx.GetTransaction().FromAddress),
						zap.Any("reward", reward),
					)
				default:
					ctx.Error(
						"resultCh is full, dropping onboarding reward",
						zap.String("user", tx.GetTransaction().FromAddress),
					)
				}
			}

			// 累積交易量以便處理 Share Pool Task
			if task.Type == model.TaskType_TASK_TYPE_SHARE_POOL {
				// TODO: 2024/11/25|sean|!! fix me !! you need to get usdc amount from swap details
				userSwapVolume[tx.GetTransaction().FromAddress] += float64(tx.GetTransaction().Amount)
				totalSwapVolume += float64(tx.GetTransaction().Amount)
			}
		}
	}

	// 4. 分配 Share Pool Task 獎勵
	for _, task := range campaign.Tasks() {
		if task.Type == model.TaskType_TASK_TYPE_SHARE_POOL {
			for user, volume := range userSwapVolume {
				// 檢查用戶是否完成 Onboarding Task
				if !campaign.HasCompletedOnboardingTask(volume) {
					continue
				}

				// 計算用戶獎勵
				points := int64((volume / totalSwapVolume) * 10000) // 假設總分數為 10,000

				// 創建獎勵
				reward := &model.Reward{
					Id:         "", // 生成唯一 ID
					UserId:     user,
					CampaignId: campaign.Id,
					Points:     points,
				}

				// 發送到結果通道
				select {
				case resultCh <- reward:
					ctx.Info("Share Pool Task reward sent", zap.String("user", user), zap.Any("reward", reward))
				default:
					ctx.Error("resultCh is full, dropping share pool reward", zap.String("user", user))
				}
			}
		}
	}

	return nil
}
