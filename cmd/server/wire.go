//go:build wireinject

//go:generate wire

package server

import (
	"github.com/blackhorseya/pelith-assessment/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewCmd(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		newImpl,
	))
}
