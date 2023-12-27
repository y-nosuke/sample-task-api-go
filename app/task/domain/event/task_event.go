package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
)

type TaskEvent[T TaskEventData] interface {
	event.DomainEvent
	TaskID() uuid.UUID
	Type() string
	Data() T
	Created(CreatedBy *uuid.UUID, CreatedAt *time.Time)
}

type TaskEventData interface {
}

type TaskEventCommon[T TaskEventData] struct {
	event.DomainEvent
	taskID    *uuid.UUID
	eventType string
	data      T
	CreatedBy *uuid.UUID
	CreatedAt *time.Time
}

func newTaskEventCommon[T TaskEventData](taskID *uuid.UUID, eventType string, data T) *TaskEventCommon[T] {
	return &TaskEventCommon[T]{
		DomainEvent: event.NewDomainEvent(),
		taskID:      taskID,
		eventType:   eventType,
		data:        data,
	}
}

func (t *TaskEventCommon[T]) TaskID() uuid.UUID {
	return *t.taskID
}

func (t *TaskEventCommon[T]) Type() string {
	return t.eventType
}

func (t *TaskEventCommon[T]) Created(CreatedBy *uuid.UUID, CreatedAt *time.Time) {
	t.CreatedBy = CreatedBy
	t.CreatedAt = CreatedAt
}

func (t *TaskEventCommon[T]) Data() T {
	return t.data
}
