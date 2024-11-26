//go:build wireinject

//go:generate wire

package server

import (
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/composite"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/external/etherscan"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/storage/pg"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/transports/grpc"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/transports/http"
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
	return config.GetService(serviceName)
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
		grpc.NewInitServersFn,
		grpc.NewHealthServer,
		grpc.NewCampaignServer,

		// app layer
		command.NewCreateCampaignHandler,
		command.NewAddTaskHandler,
		command.NewStartCampaignHandler,
		command.NewRunBacktestHandler,
		query.NewTaskQueryService,
		query.NewTransactionQueryService,

		// entity layer
		biz.NewCampaignService,
		biz.NewTaskService,
		biz.NewBacktestService,

		// repo layer
		pg.NewCampaignRepo,
		pg.NewCampaignCreator,
		pg.NewCampaignGetter,
		pg.NewTaskRepo,
		pg.NewTaskCreator,
		pg.NewTaskGetter,
		pg.NewTransactionRepoImpl,
		pg.NewCampaignUpdater,
		etherscan.NewTransactionRepoImpl,
		etherscan.NewTransactionGetter,
		composite.NewTransactionCompositeRepoImpl,

		// infra
		httpx.NewGinServer,
		grpcx.NewServer,
		pgx.NewClient,
	))
}
