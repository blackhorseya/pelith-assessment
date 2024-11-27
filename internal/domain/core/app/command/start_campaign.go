package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

// CampaignStrategy is the interface for campaign strategies.
type CampaignStrategy interface {
	Execute(c context.Context, campaign *biz.Campaign) error
}

type emptyStrategy struct{}

func (s *emptyStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	return nil
}

type backtestStrategy struct {
	backtestService biz.BacktestService
	campaignUpdater CampaignUpdater
}

func (s *backtestStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	ctx := contextx.WithContext(c)

	err := campaign.Start()
	if err != nil {
		ctx.Error("failed to start campaign", zap.Error(err), zap.Any("campaign", &campaign))
		return err
	}

	rewardCh := make(chan *model.Reward)
	go func() {
		err = s.backtestService.RunBacktest(context.Background(), campaign, rewardCh)
		if err != nil {
			ctx.Error("failed to run backtest", zap.Error(err))
		}
		close(rewardCh)
	}()

	go func() {
		for reward := range rewardCh {
			err = s.campaignUpdater.DistributeReward(context.Background(), reward)
			if err != nil {
				ctx.Error("failed to distribute reward", zap.Error(err))
				continue
			}
		}
	}()

	return nil
}

type realTimeStrategy struct {
	transactionAdapter query.TransactionAdapter
	campaignUpdater    CampaignUpdater
}

func (s *realTimeStrategy) Execute(c context.Context, campaign *biz.Campaign) error {
	ctx := contextx.WithContext(c)

	err := campaign.Start()
	if err != nil {
		ctx.Error("failed to start campaign", zap.Error(err), zap.Any("campaign", &campaign))
		return err
	}

	txCh := make(chan *biz.Transaction)
	go func() {
		err = s.transactionAdapter.GetSwapTxByPoolAddress(context.Background(), campaign.PoolId, txCh)
		if err != nil {
			ctx.Error("failed to get swapTx by pool address", zap.Error(err))
		}
		close(txCh)
	}()

	go func() {
		for tx := range txCh {
			reward, err2 := campaign.OnSwapExecuted(tx)
			if err2 != nil {
				ctx.Error("failed to handle swap executed", zap.Error(err2))
				continue
			}
			if reward == nil {
				continue
			}

			err = s.campaignUpdater.DistributeReward(context.Background(), reward)
			if err != nil {
				ctx.Error("failed to distribute reward", zap.Error(err))
				continue
			}
		}
	}()

	return nil
}

// StartCampaignHandler is the handler for starting a campaign.
type StartCampaignHandler struct {
	strategies      map[model.CampaignMode]CampaignStrategy
	campaignGetter  query.CampaignGetter
	campaignUpdater CampaignUpdater
}

// NewStartCampaignHandler creates a new StartCampaignHandler instance.
func NewStartCampaignHandler(
	campaignGetter query.CampaignGetter,
	campaignUpdater CampaignUpdater,
	backtestService biz.BacktestService,
	transactionAdapter query.TransactionAdapter,
) *StartCampaignHandler {
	return &StartCampaignHandler{
		strategies: map[model.CampaignMode]CampaignStrategy{
			model.CampaignMode_CAMPAIGN_MODE_UNSPECIFIED: &emptyStrategy{},
			model.CampaignMode_CAMPAIGN_MODE_BACKTEST: &backtestStrategy{
				backtestService: backtestService,
			},
			model.CampaignMode_CAMPAIGN_MODE_REAL_TIME: &realTimeStrategy{
				transactionAdapter: transactionAdapter,
				campaignUpdater:    campaignUpdater,
			},
		},
		campaignGetter:  campaignGetter,
		campaignUpdater: campaignUpdater,
	}
}

func (h StartCampaignHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	ctx := contextx.WithContext(c)

	cmd, ok := msg.(StartCampaignCommand)
	if !ok {
		ctx.Error("invalid command type for StartCampaignHandler", zap.Any("command", &msg))
		return "", errors.New("invalid command type for StartCampaignHandler")
	}

	campaign, err := h.campaignGetter.GetByID(ctx, cmd.ID)
	if err != nil {
		ctx.Error("failed to fetch campaign", zap.Error(err))
		return "", err
	}
	if campaign == nil {
		return "", errors.New("campaign not found")
	}

	strategy, ok := h.strategies[campaign.Mode]
	if !ok {
		ctx.Error("strategy not found", zap.Any("mode", campaign.Mode))
		return "", errors.New("strategy not found")
	}

	err = strategy.Execute(ctx, campaign)
	if err != nil {
		ctx.Error("failed to execute strategy", zap.Error(err))
		return "", err
	}

	err = h.campaignUpdater.UpdateStatus(ctx, campaign, model.CampaignStatus_CAMPAIGN_STATUS_ACTIVE)
	if err != nil {
		ctx.Error("failed to update campaign status", zap.Error(err))
		return "", err
	}

	return campaign.Id, nil
}
