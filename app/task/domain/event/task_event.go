package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
)

type TaskEvent struct {
	event.DomainEvent
	TaskID     uuid.UUID
	OccurredBy uuid.UUID
	OccurredAt time.Time
}

func newTaskEvent(taskID uuid.UUID, occurredBy uuid.UUID, occurredAt time.Time) *TaskEvent {
	return &TaskEvent{
		DomainEvent: event.NewDomainEvent(),
		TaskID:      taskID,
		OccurredBy:  occurredBy,
		OccurredAt:  occurredAt,
	}
}
