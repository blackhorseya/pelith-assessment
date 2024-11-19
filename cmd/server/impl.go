package server

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/pkg/adapterx"
)

type impl struct {
	injector  *injector
	ginServer *httpx.GinServer
}

func newImpl(injector *injector, ginServer *httpx.GinServer) adapterx.Server {
	return &impl{
		injector:  injector,
		ginServer: ginServer,
	}
}

func (i *impl) Start(c context.Context) error {
	// TODO: 2024/11/20|sean|implement me
	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	// TODO: 2024/11/20|sean|implement me
	return nil
}
