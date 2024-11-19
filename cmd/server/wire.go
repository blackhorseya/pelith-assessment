//go:build wireinject

//go:generate wire

package server

import (
	"github.com/blackhorseya/pelith-assessment/cmd/server/wirex"
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
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
		wire.Struct(new(wirex.Injector), "*"),
		initConfigx,
		initAPP,

		// adapter
		NewInitRoutesFn,

		// infra
		httpx.NewGinServer,
	))
}
