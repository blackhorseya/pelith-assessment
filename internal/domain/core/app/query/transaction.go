//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"strings"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// ListTransactionCondition is the condition for list transaction.
type ListTransactionCondition struct {
	ContractAddress string
	StartTime       time.Time
	EndTime         time.Time
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

	// 從 TransactionGetter 查詢交易數據
	transactions, _, err := s.txGetter.ListByAddress(ctx, address, ListTransactionCondition{
		// TODO: 2024/11/22|sean|hard code contract address
		ContractAddress: "0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc",
		StartTime:       campaign.StartTime.AsTime(),
		EndTime:         campaign.EndTime.AsTime(),
	})
	if err != nil {
		ctx.Error("failed to fetch transactions", zap.Error(err))
		return 0, err
	}

	// 計算總數量
	var totalAmount float64
	for _, tx := range transactions {
		for _, task := range campaign.Tasks {
			if strings.EqualFold(task.Criteria.PoolId, tx.ToAddress) {
				totalAmount += float64(tx.Amount)
			}
		}
	}

	return totalAmount, nil
}
