package pg

import (
	"time"
)

// TransactionDAO represents the transactions table
type TransactionDAO struct {
	TxHash      string    `db:"tx_hash"`      // TransactionDAO hash
	BlockNumber int64     `db:"block_number"` // Block number
	Timestamp   time.Time `db:"timestamp"`    // TransactionDAO timestamp
	FromAddress string    `db:"from_address"` // Sender address
	ToAddress   string    `db:"to_address"`   // Receiver address
	Value       int64     `db:"value"`        // TransactionDAO value
	GasUsed     int       `db:"gas_used"`     // Gas used
	GasPrice    float64   `db:"gas_price"`    // Gas price
	Status      bool      `db:"status"`       // TransactionDAO status (success or failure)
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
