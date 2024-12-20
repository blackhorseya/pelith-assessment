//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

// ListTransactionCondition is the condition for list transaction.
type ListTransactionCondition struct {
	PoolAddress string
	StartTime   time.Time
	EndTime     time.Time
}

// GetLogsCondition is the condition for get logs.
type GetLogsCondition struct {
	StartTime time.Time
	EndTime   time.Time
}

// TransactionRepo is the adapter for transaction.
type TransactionRepo interface {
	GetSwapTxByPoolAddress(
		c context.Context,
		address string,
		cond ListTransactionCondition,
		txCh chan<- *biz.Transaction,
	) error

	GetSwapTxByUserAddressAndPoolAddress(
		c context.Context,
		address, poolAddress string,
		cond ListTransactionCondition,
		txCh chan<- *biz.Transaction,
	) error
}

// TransactionAdapter is the adapter for transaction.
type TransactionAdapter interface {
	GetSwapTxByPoolAddress(
		c context.Context,
		contractAddress string,
		txCh chan<- *biz.Transaction,
	) error
}

// // TransactionQueryService is the service for transaction query.
// type TransactionQueryService struct {
// 	txGetter       TransactionGetter
// 	campaignGetter CampaignGetter
// }

// // GetTotalSwapAmount calculates the total swap amount for a given address and campaign ID.
// func (s *TransactionQueryService) GetTotalSwapAmount(c context.Context, address,
// campaignID string) (float64, error) {
// 	ctx := contextx.WithContext(c)
//
// 	// Retrieve the campaign by ID
// 	campaign, err := s.campaignGetter.GetByID(ctx, campaignID)
// 	if err != nil {
// 		ctx.Error("failed to fetch campaign", zap.Error(err), zap.String("campaign_id", campaignID))
// 		return 0, fmt.Errorf("failed to fetch campaign: %w", err)
// 	}
//
// 	if campaign == nil || len(campaign.Tasks()) == 0 {
// 		ctx.Warn("campaign has no tasks", zap.String("campaign_id", campaignID))
// 		return 0, nil
// 	}
//
// 	// Fetch transactions for the specified address and campaign's criteria
// 	transactions, _, err := s.txGetter.ListByAddress(ctx, address, ListTransactionCondition{
// 		PoolAddress: campaign.PoolId,
// 		StartTime:   campaign.StartTime.AsTime(),
// 		EndTime:     campaign.EndTime.AsTime(),
// 	})
// 	if err != nil {
// 		ctx.Error("failed to fetch transactions", zap.Error(err))
// 		return 0, fmt.Errorf("failed to fetch transactions: %w", err)
// 	}
//
// 	// Compute the total amount of USDC swapped
// 	totalAmount, err := calculateTotalUSDC(ctx, transactions, usdcAddress)
// 	if err != nil {
// 		return 0, fmt.Errorf("error calculating total USDC: %w", err)
// 	}
//
// 	return totalAmount, nil
// }
//
// // calculateTotalUSDC computes the total amount of USDC from swap transactions.
// func calculateTotalUSDC(ctx contextx.Contextx, transactions []*biz.Transaction,
// usdcAddress string) (float64, error) {
// 	var totalAmount float64
//
// 	for _, tx := range transactions {
// 		// Skip non-swap transactions
// 		if tx.IsSwapType() {
// 			continue
// 		}
//
// 		for _, detail := range tx.GetTransaction().SwapDetails {
// 			// Process USDC "from" and "to" amounts
// 			if amount, err := getUSDCAmount(detail.FromTokenAddress, detail.FromTokenAmount, usdcAddress); err != nil {
// 				ctx.Error("failed to parse FromTokenAmount", zap.Error(err))
// 				return 0, err
// 			} else {
// 				totalAmount += amount
// 			}
//
// 			if amount, err := getUSDCAmount(detail.ToTokenAddress, detail.ToTokenAmount, usdcAddress); err != nil {
// 				ctx.Error("failed to parse ToTokenAmount", zap.Error(err))
// 				return 0, err
// 			} else {
// 				totalAmount += amount
// 			}
// 		}
// 	}
//
// 	return totalAmount, nil
// }
//
// // getUSDCAmount parses the token amount if the token address matches the USDC address.
// func getUSDCAmount(tokenAddress, tokenAmount, usdcAddress string) (float64, error) {
// 	if strings.EqualFold(tokenAddress, usdcAddress) {
// 		return strconv.ParseFloat(tokenAmount, 64)
// 	}
// 	return 0, nil
// }
