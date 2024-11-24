package eventx

import (
	"time"
)

// DomainEvent represents a generic event with basic properties.
type DomainEvent interface {
	GetOccurredAt() time.Time
	GetName() string
	GetVersion() int
	GetPayload() interface{}
}

// EventHandler is the interface for handling domain events.
type EventHandler interface {
	Handle(event DomainEvent)
}

// EventBus is the interface for a pluggable event bus.
type EventBus interface {
	Subscribe(handler EventHandler) error
	Publish(event DomainEvent) error
}
