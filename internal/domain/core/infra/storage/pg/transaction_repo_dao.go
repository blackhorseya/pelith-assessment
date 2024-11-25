package pg

import (
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// TransactionDAO represents the transactions table
type TransactionDAO struct {
	TxHash      string    `db:"tx_hash"`      // TransactionDAO hash
	BlockNumber int64     `db:"block_number"` // Block number
	Timestamp   time.Time `db:"timestamp"`    // TransactionDAO timestamp
	FromAddress string    `db:"from_address"` // Sender address
	ToAddress   string    `db:"to_address"`   // Receiver address
}

// FromBizTransactionToDAO converts a biz.Transaction to a TransactionDAO
func FromBizTransactionToDAO(transaction *biz.Transaction) *TransactionDAO {
	return &TransactionDAO{
		TxHash:      transaction.TxHash,
		BlockNumber: transaction.BlockNumber,
		Timestamp:   transaction.Timestamp.AsTime(),
		FromAddress: transaction.FromAddress,
		ToAddress:   transaction.ToAddress,
	}
}

// SwapEventDAO represents the swap_events table
type SwapEventDAO struct {
	ID               int    `db:"id"`                 // Primary key
	TxHash           string `db:"tx_hash"`            // Associated transaction hash
	FromTokenAddress string `db:"from_token_address"` // Source token address
	ToTokenAddress   string `db:"to_token_address"`   // Destination token address
	FromTokenAmount  string `db:"from_token_amount"`  // Source token amount
	ToTokenAmount    string `db:"to_token_amount"`    // Destination token amount
	PoolAddress      string `db:"pool_address"`       // Swap pool address (if applicable)
}

// FromModelSwapDetailToDAO converts a model.SwapDetail to a SwapEventDAO
func FromModelSwapDetailToDAO(txHash string, swap *model.SwapDetail) *SwapEventDAO {
	return &SwapEventDAO{
		TxHash:           txHash,
		FromTokenAddress: swap.FromTokenAddress,
		ToTokenAddress:   swap.ToTokenAddress,
		FromTokenAmount:  swap.FromTokenAmount,
		ToTokenAmount:    swap.ToTokenAmount,
		PoolAddress:      swap.PoolAddress,
	}
}
