package biz

import (
	"context"
	"strconv"

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

func (i *backtestServiceImpl) RunBacktest(
	c context.Context,
	campaign *biz.Campaign,
	resultCh chan<- *model.Reward,
) error {
	ctx := contextx.WithContext(c)

	// 1. 獲取交易記錄
	txCh := make(chan *biz.Transaction)
	var err error
	go func() {
		err = i.txRepo.GetSwapTxByPoolAddress(c, campaign.PoolId, query.ListTransactionCondition{
			StartTime: campaign.StartTime.AsTime(),
			EndTime:   campaign.EndTime.AsTime(),
		}, txCh)
		if err != nil {
			ctx.Error("failed to get swapTx swapTx by pool address", zap.Error(err))
		}
		close(txCh)
	}()

	// 2. 準備累積數據
	const usdcAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	userSwapVolume := make(map[string]float64) // 用戶的交易量 usdc
	userOnboardingReward := make(map[string]bool)
	totalSwapVolume := 0.0 // 總交易量
	onboardingTask := campaign.GetTaskByType(model.TaskType_TASK_TYPE_ONBOARDING)

	// 3. 處理交易記錄
	for swapTx := range txCh {
		amount, err2 := strconv.ParseFloat(swapTx.GetSwapAmountByTokenAddress(usdcAddress), 64)
		if err2 != nil {
			ctx.Error("failed to parse swap amount", zap.Error(err2))
			continue
		}
		userSwapVolume[swapTx.GetTransaction().FromAddress] += amount
		totalSwapVolume += amount

		if !userOnboardingReward[swapTx.GetTransaction().FromAddress] {
			totalAmount := userSwapVolume[swapTx.GetTransaction().FromAddress]
			if totalAmount >= onboardingTask.Criteria.MinTransactionAmount {
				userOnboardingReward[swapTx.GetTransaction().FromAddress] = true
				reward := &model.Reward{
					Id:         "", // 生成唯一 ID
					UserId:     swapTx.GetTransaction().FromAddress,
					CampaignId: campaign.Id,
					Points:     100, // 固定獎勵點數
				}

				// 發送到結果通道
				select {
				case resultCh <- reward:
					ctx.Info(
						"Onboarding Task reward sent",
						zap.String("user", swapTx.GetTransaction().FromAddress),
						zap.Any("reward", &reward),
					)
				default:
					ctx.Error(
						"resultCh is full, dropping onboarding reward",
						zap.String("user", swapTx.GetTransaction().FromAddress),
					)
				}
			}
		}
	}

	// 4. 分配 Share Pool Task 獎勵
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
			ctx.Info("Share Pool Task reward sent", zap.String("user", user), zap.Any("reward", &reward))
		default:
			ctx.Error("resultCh is full, dropping share pool reward", zap.String("user", user))
		}
	}

	return nil
}
