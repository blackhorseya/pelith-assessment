package adapterx

import (
	"context"
)

// Server is an interface that represents the daemon.
type Server interface {
	// Start is used to start the daemon.
	Start(c context.Context) error

	// Shutdown is used to shut down the daemon.
	Shutdown(c context.Context) error
}
