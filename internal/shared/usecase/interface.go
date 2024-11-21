//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package usecase

import (
	"context"
)

type (
	// Handler is the interface for handling a message.
	Handler interface {
		Handle(context.Context, Message) (string, error)
	}

	// PipelineBehaviour is the interface for processing a message in a pipeline.
	PipelineBehaviour interface {
		Process(context.Context, Message, Next) error
	}

	// Message is the interface for a message.
	Message interface {
		Key() int
	}
)
