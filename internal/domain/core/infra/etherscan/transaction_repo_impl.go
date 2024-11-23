package etherscan

import (
	"context"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TransactionRepoImpl is the implementation of TransactionRepo.
type TransactionRepoImpl struct {
	etherscanAPI *etherscan.Client
	ethclientAPI *ethclient.Client
}

// NewTransactionRepoImpl is used to create a new TransactionRepoImpl.
func NewTransactionRepoImpl(app *configx.Application) (*TransactionRepoImpl, error) {
	etherscanAPI := etherscan.New(etherscan.Mainnet, app.Etherscan.APIKey)

	ethclientAPI, err := ethclient.Dial("https://mainnet.infura.io/v3/" + app.Infura.ProjectID)
	if err != nil {
		return nil, err
	}

	return &TransactionRepoImpl{
		etherscanAPI: etherscanAPI,
		ethclientAPI: ethclientAPI,
	}, nil
}

// NewTransactionGetter is used to create a new TransactionGetter.
func NewTransactionGetter(impl *TransactionRepoImpl) query.TransactionGetter {
	return impl
}

func (i *TransactionRepoImpl) ListByAddress(
	c context.Context,
	address string,
	cond query.ListTransactionCondition,
) (item biz.TransactionList, total int, err error) {
	ctx := contextx.WithContext(c)

	startBlock, err := i.etherscanAPI.BlockNumber(cond.StartTime.Unix(), "after")
	if err != nil {
		ctx.Error("failed to fetch start block", zap.Error(err), zap.Time("start_time", cond.StartTime))
		return nil, 0, err
	}

	// if cond.EndTime > now then set endBlock = now
	if cond.EndTime.After(time.Now()) {
		cond.EndTime = time.Now()
	}
	endBlock, err := i.etherscanAPI.BlockNumber(cond.EndTime.Unix(), "before")
	if err != nil {
		ctx.Error("failed to fetch end block", zap.Error(err), zap.Time("end_time", cond.EndTime))
		return nil, 0, err
	}

	txs, err := i.etherscanAPI.NormalTxByAddress(address, &startBlock, &endBlock, 1, 100, true)
	if err != nil {
		ctx.Error("failed to fetch transactions", zap.Error(err))
		return nil, 0, err
	}

	var res biz.TransactionList
	for _, tx := range txs {
		res = append(res, &biz.Transaction{
			Transaction: model.Transaction{
				TxHash:      tx.Hash,
				FromAddress: tx.From,
				ToAddress:   tx.To,
				Amount:      tx.Value.Int().Int64(),
				Timestamp:   timestamppb.New(tx.TimeStamp.Time()),
				TaskId:      nil,
				CampaignId:  nil,
				Status:      model.TransactionStatus_TRANSACTION_STATUS_COMPLETED,
				Type:        0,
				SwapDetails: &model.SwapDetails{
					FromTokenAddress: "",
					ToTokenAddress:   "",
					FromTokenAmount:  0,
					ToTokenAmount:    0,
					PoolAddress:      "",
				},
			},
		})
	}

	return res, len(res), nil
}
