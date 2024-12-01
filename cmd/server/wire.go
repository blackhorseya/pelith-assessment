//go:build wireinject

//go:generate wire

package server

import (
	"os"

	command2 "github.com/blackhorseya/pelith-assessment/internal/domain/app/command"
	query2 "github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	biz2 "github.com/blackhorseya/pelith-assessment/internal/domain/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/composite"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/external/etherscan"
	pg2 "github.com/blackhorseya/pelith-assessment/internal/domain/infra/storage/pg"
	grpc2 "github.com/blackhorseya/pelith-assessment/internal/domain/infra/transports/grpc"
	http2 "github.com/blackhorseya/pelith-assessment/internal/domain/infra/transports/http"
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/pgx"
	"github.com/blackhorseya/pelith-assessment/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const serviceName = "server"

func initConfigx(v *viper.Viper) (*configx.Configx, error) {
	return configx.LoadConfig(v.GetString("config"))
}

func initAPP(config *configx.Configx) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, err
	}

	if app.Etherscan.APIKey == "" {
		app.Etherscan.APIKey = os.Getenv("ETHERSCAN_API_KEY")
	}

	if app.Infura.ProjectID == "" {
		app.Infura.ProjectID = os.Getenv("INFURA_PROJECT_ID")
	}

	return app, nil
}

func NewCmd(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		newImpl,
		wire.Struct(new(Injector), "*"),
		initConfigx,
		initAPP,

		// adapter
		http2.NewInitUserRoutesFn,
		http2.NewQueryController,
		http2.NewCommandController,
		grpc2.NewInitServersFn,
		grpc2.NewHealthServer,
		grpc2.NewCampaignServer,

		// app layer
		command2.NewCreateCampaignHandler,
		command2.NewAddTaskHandler,
		command2.NewStartCampaignHandler,
		command2.NewRunBacktestHandler,
		query2.NewRewardQueryStore,
		query2.NewUserQueryStore,

		// entity layer
		biz2.NewCampaignService,
		grpc2.NewCampaignServiceClient,
		biz2.NewTaskService,
		biz2.NewBacktestService,
		biz2.NewUserService,

		// repo layer
		pg2.NewCampaignRepo,
		pg2.NewCampaignCreator,
		pg2.NewCampaignGetter,
		pg2.NewCampaignUpdater,
		pg2.NewCampaignDeleter,
		pg2.NewTaskRepo,
		pg2.NewTaskCreator,
		pg2.NewTransactionRepoImpl,
		pg2.NewRewardRepo,
		pg2.NewRewardGetter,
		etherscan.NewTransactionRepoImpl,
		etherscan.NewTransactionAdapter,
		composite.NewTransactionCompositeRepoImpl,
		composite.NewTransactionRepoImpl,

		// infra
		httpx.NewGinServer,
		grpcx.NewServer,
		grpcx.NewClient,
		pgx.NewClient,
	))
}
