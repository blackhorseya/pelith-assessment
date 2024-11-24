package biz

import (
	"time"

	"github.com/blackhorseya/pelith-assessment/pkg/eventx"
	"github.com/google/uuid"
)

var _ eventx.DomainEvent = (*SwapExecutedEvent)(nil) // 確保實現接口

// SwapExecutedPayload is the payload for swap executed.
type SwapExecutedPayload struct {
	TxID string
}

// SwapExecutedEvent is the event for swap executed.
type SwapExecutedEvent struct {
	id         string
	occurredAt time.Time
	payload    SwapExecutedPayload
}

// NewSwapExecutedEvent creates a new SwapExecutedEvent instance.
func NewSwapExecutedEvent(occurredAt time.Time, payload SwapExecutedPayload) *SwapExecutedEvent {
	return &SwapExecutedEvent{
		id:         uuid.New().String(),
		occurredAt: occurredAt,
		payload:    payload,
	}
}

func (e *SwapExecutedEvent) GetID() string {
	return e.id
}

func (e *SwapExecutedEvent) GetOccurredAt() time.Time {
	return e.occurredAt
}

func (e *SwapExecutedEvent) GetName() string {
	return "SwapExecuted"
}

func (e *SwapExecutedEvent) GetPayload() interface{} {
	return e.payload
}
