package server

import (
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
)

type injector struct {
	C *configx.Configx
	A *configx.Application
}
