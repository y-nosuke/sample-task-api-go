package event

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
)

type TaskEvent struct {
	event.DomainEvent
	TaskID     uuid.UUID
	OccurredBy uuid.UUID
	OccurredAt time.Time
}

func newTaskEvent(taskID uuid.UUID, occurredBy uuid.UUID, occurredAt time.Time) (*TaskEvent, error) {
	domainEvent, err := event.NewDomainEvent()
	if err != nil {
		return nil, xerrors.Errorf("uuid.NewV7(): %w", err)
	}
	return &TaskEvent{
		DomainEvent: domainEvent,
		TaskID:      taskID,
		OccurredBy:  occurredBy,
		OccurredAt:  occurredAt,
	}, nil
}
