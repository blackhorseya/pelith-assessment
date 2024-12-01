//go:build wireinject

//go:generate wire

package server

import (
	"os"

	"github.com/blackhorseya/pelith-assessment/internal/domain/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/composite"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/external/etherscan"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/storage/mongodb"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/storage/pg"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/transports/grpc"
	"github.com/blackhorseya/pelith-assessment/internal/domain/infra/transports/http"
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/mongodbx"
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
		http.NewInitUserRoutesFn,
		http.NewQueryController,
		http.NewCommandController,
		grpc.NewInitServersFn,
		grpc.NewHealthServer,
		grpc.NewCampaignServer,

		// app layer
		command.NewCreateCampaignHandler,
		command.NewAddTaskHandler,
		command.NewStartCampaignHandler,
		command.NewRunBacktestHandler,
		query.NewRewardQueryStore,
		query.NewUserQueryStore,

		// entity layer
		biz.NewCampaignService,
		biz.NewTaskService,
		biz.NewBacktestService,
		biz.NewUserService,
		grpc.NewCampaignServiceClient,

		// repo layer
		pg.NewCampaignRepo,
		pg.NewCampaignGetter,
		pg.NewCampaignUpdater,
		pg.NewCampaignDeleter,
		pg.NewTaskRepo,
		pg.NewTaskCreator,
		pg.NewTransactionRepoImpl,
		pg.NewRewardRepo,
		pg.NewRewardGetter,
		etherscan.NewTransactionRepoImpl,
		etherscan.NewTransactionAdapter,
		composite.NewTransactionCompositeRepoImpl,
		composite.NewTransactionRepoImpl,
		mongodb.NewCampaignCreator,
		mongodb.NewCampaignRepoImpl,

		// infra
		httpx.NewGinServer,
		grpcx.NewServer,
		grpcx.NewClient,
		pgx.NewClient,
		mongodbx.NewClient,
	))
}
