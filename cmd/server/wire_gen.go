// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/external/etherscan"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/storage/pg"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/transports/grpc"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/transports/http"
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/pgx"
	"github.com/blackhorseya/pelith-assessment/pkg/adapterx"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewCmd(v *viper.Viper) (adapterx.Server, func(), error) {
	configx, err := initConfigx(v)
	if err != nil {
		return nil, nil, err
	}
	application, err := initAPP(configx)
	if err != nil {
		return nil, nil, err
	}
	injector := &Injector{
		C: configx,
		A: application,
	}
	db, err := pgx.NewClient(application)
	if err != nil {
		return nil, nil, err
	}
	taskRepoImpl := pg.NewTaskRepo(db)
	taskGetter := pg.NewTaskGetter(taskRepoImpl)
	transactionRepoImpl, err := etherscan.NewTransactionRepoImpl(application)
	if err != nil {
		return nil, nil, err
	}
	transactionGetter := etherscan.NewTransactionGetter(transactionRepoImpl)
	campaignRepoImpl, err := pg.NewCampaignRepo(db)
	if err != nil {
		return nil, nil, err
	}
	campaignGetter, err := pg.NewCampaignGetter(campaignRepoImpl)
	if err != nil {
		return nil, nil, err
	}
	transactionQueryService := query.NewTransactionQueryService(transactionGetter, campaignGetter)
	taskQueryService := query.NewTaskQueryService(taskGetter, transactionQueryService)
	queryController := http.NewQueryController(taskQueryService)
	initRoutes := http.NewInitUserRoutesFn(queryController)
	ginServer, err := httpx.NewGinServer(application, initRoutes)
	if err != nil {
		return nil, nil, err
	}
	campaignService := biz.NewCampaignService()
	campaignCreator, err := pg.NewCampaignCreator(campaignRepoImpl)
	if err != nil {
		return nil, nil, err
	}
	createCampaignHandler := command.NewCreateCampaignHandler(campaignService, campaignCreator)
	taskService := biz.NewTaskService()
	taskCreator := pg.NewTaskCreator(taskRepoImpl)
	addTaskHandler := command.NewAddTaskHandler(campaignService, campaignGetter, taskService, taskCreator)
	backtestService := biz.NewBacktestService(transactionGetter)
	startCampaignHandler := command.NewStartCampaignHandler(campaignGetter, backtestService)
	campaignServiceServer := grpc.NewCampaignServer(createCampaignHandler, addTaskHandler, startCampaignHandler, campaignGetter)
	initServers := grpc.NewInitServersFn(campaignServiceServer)
	healthServer := grpc.NewHealthServer()
	server, err := grpcx.NewServer(application, initServers, healthServer)
	if err != nil {
		return nil, nil, err
	}
	adapterxServer := newImpl(injector, ginServer, server)
	return adapterxServer, func() {
	}, nil
}

// wire.go:

const serviceName = "server"

func initConfigx(v *viper.Viper) (*configx.Configx, error) {
	return configx.LoadConfig(v.GetString("config"))
}

func initAPP(config *configx.Configx) (*configx.Application, error) {
	return config.GetService(serviceName)
}
