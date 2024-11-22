//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

// TransactionGetter is used to get the transaction.
type TransactionGetter interface {
}

// TransactionQueryService is the service for transaction query.
type TransactionQueryService struct {
	txGetter TransactionGetter
}

// NewTransactionQueryService is used to create a new TransactionQueryService.
func NewTransactionQueryService(txGetter TransactionGetter) *TransactionQueryService {
	return &TransactionQueryService{txGetter: txGetter}
}
