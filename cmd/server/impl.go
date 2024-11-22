package server

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/pkg/adapterx"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	injector   *Injector
	ginServer  *httpx.GinServer
	grpcserver *grpcx.Server
}

func newImpl(injector *Injector, ginServer *httpx.GinServer, grpcserver *grpcx.Server) adapterx.Server {
	return &impl{
		injector:   injector,
		ginServer:  ginServer,
		grpcserver: grpcserver,
	}
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)

	err := i.ginServer.Start(ctx)
	if err != nil {
		ctx.Error("gin server start failed", zap.Error(err))
		return err
	}

	err = i.grpcserver.Start(ctx)
	if err != nil {
		ctx.Error("grpc server start failed", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)

	err := i.ginServer.Stop(ctx)
	if err != nil {
		ctx.Error("gin server stop failed", zap.Error(err))
	}

	err = i.grpcserver.Stop(ctx)
	if err != nil {
		ctx.Error("grpc server stop failed", zap.Error(err))
	}

	return nil
}
