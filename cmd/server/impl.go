package server

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/pkg/adapterx"
)

type impl struct {
}

func newImpl() adapterx.Server {
	return &impl{}
}

func (i *impl) Start(c context.Context) error {
	// TODO: 2024/11/20|sean|implement me
	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	// TODO: 2024/11/20|sean|implement me
	return nil
}
