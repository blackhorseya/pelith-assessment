package biz

import (
	"time"

	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
)

var _ eventx.DomainEvent = (*SwapExecutedEvent)(nil) // 確保實現接口

// SwapExecutedEvent is the event for swap executed.
type SwapExecutedEvent struct {
	occurredAt time.Time
	version    int

	payload interface{}
}

// NewSwapExecutedEvent creates a new SwapExecutedEvent instance.
func NewSwapExecutedEvent(occurredAt time.Time, version int, payload interface{}) *SwapExecutedEvent {
	return &SwapExecutedEvent{
		occurredAt: occurredAt,
		version:    version,
		payload:    payload,
	}
}

func (e *SwapExecutedEvent) GetOccurredAt() time.Time {
	return e.occurredAt
}

func (e *SwapExecutedEvent) GetName() string {
	return "SwapExecuted"
}

func (e *SwapExecutedEvent) GetVersion() int {
	return e.version
}

func (e *SwapExecutedEvent) GetPayload() interface{} {
	return e.payload
}
