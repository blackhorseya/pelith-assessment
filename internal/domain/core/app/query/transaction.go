//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// ListTransactionCondition is the condition for list transaction.
type ListTransactionCondition struct {
	PoolAddress string
	StartTime   time.Time
	EndTime     time.Time
}

// TransactionGetter is used to get the transaction.
type TransactionGetter interface {
	ListByAddress(
		c context.Context,
		address string,
		cond ListTransactionCondition,
	) (item biz.TransactionList, total int, err error)
}

// TransactionQueryService is the service for transaction query.
type TransactionQueryService struct {
	txGetter       TransactionGetter
	campaignGetter CampaignGetter
}

// NewTransactionQueryService is used to create a new TransactionQueryService.
func NewTransactionQueryService(txGetter TransactionGetter, campaignGetter CampaignGetter) *TransactionQueryService {
	return &TransactionQueryService{
		txGetter:       txGetter,
		campaignGetter: campaignGetter,
	}
}

// GetTotalSwapAmount is used to get the total swap amount.
func (s *TransactionQueryService) GetTotalSwapAmount(c context.Context, address, campaignID string) (float64, error) {
	ctx := contextx.WithContext(c)

	// 從 CampaignGetter 查詢 campaign
	campaign, err := s.campaignGetter.GetByID(ctx, campaignID)
	if err != nil || campaign == nil {
		ctx.Error("failed to fetch campaign", zap.Error(err), zap.String("campaign_id", campaignID))
		return 0, err
	}

	if len(campaign.Tasks) == 0 {
		ctx.Warn("no tasks in campaign", zap.String("campaign_id", campaignID))
		return 0, nil
	}

	// 從 TransactionGetter 查詢交易數據
	transactions, _, err := s.txGetter.ListByAddress(ctx, address, ListTransactionCondition{
		PoolAddress: campaign.Tasks[0].Criteria.PoolId,
		StartTime:   campaign.StartTime.AsTime(),
		EndTime:     campaign.EndTime.AsTime(),
	})
	if err != nil {
		ctx.Error("failed to fetch transactions", zap.Error(err))
		return 0, err
	}

	// 計算總數量
	const usdcAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	var totalAmount float64
	for _, tx := range transactions {
		if tx.Type == model.TransactionType_TRANSACTION_TYPE_SWAP {
			for _, detail := range tx.SwapDetails {
				if strings.EqualFold(detail.FromTokenAddress, usdcAddress) {
					value, err2 := strconv.ParseFloat(detail.FromTokenAmount, 64)
					if err2 != nil {
						ctx.Error("failed to parse float", zap.Error(err2))
						return 0, err2
					}
					totalAmount += value
				} else if strings.EqualFold(detail.ToTokenAddress, usdcAddress) {
					value, err2 := strconv.ParseFloat(detail.ToTokenAmount, 64)
					if err2 != nil {
						ctx.Error("failed to parse float", zap.Error(err2))
						return 0, err2
					}
					totalAmount += value
				}
			}
		}
	}

	return totalAmount, nil
}
