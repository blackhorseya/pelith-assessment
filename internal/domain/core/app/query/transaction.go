//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// ListTransactionCondition is the condition for list transaction.
type ListTransactionCondition struct {
	StartTime time.Time
	EndTime   time.Time
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

// GetTotalSwapUSDC 計算指定 address 和 campaignID 的 USDC 交易總數
func (s *TransactionQueryService) GetTotalSwapUSDC(c context.Context, address, campaignID string) (float64, error) {
	ctx := contextx.WithContext(c)

	// 從 TransactionGetter 查詢交易數據
	transactions, _, err := s.txGetter.ListByAddress(ctx, address, ListTransactionCondition{
		// TODO: 2024/11/22|sean|這裡要改成從 campaignID 取得對應的 StartTime 和 EndTime
		StartTime: time.Time{},
		EndTime:   time.Time{},
	})
	if err != nil {
		ctx.Error("failed to fetch transactions", zap.Error(err))
		return 0, err
	}

	// 計算總數量
	var totalUSDC float64
	for _, tx := range transactions {
		if tx.Type == model.TransactionType_TRANSACTION_TYPE_SWAP && tx.SwapDetails != nil {
			// TODO: 2024/11/22|sean|這裡要改成從 campaignID 取得對應的 USDC address
		}
	}

	return totalUSDC, nil
}
