package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"time"
)

const (
	ETaskCreated     = "TaskCreated"
	ETaskUpdated     = "TaskUpdated"
	ETaskDeleted     = "TaskDeleted"
	ETaskCompleted   = "TaskCompleted"
	ETaskUnCompleted = "TaskUnCompleted"
)

type TaskEvent interface {
	event.DomainEvent
	TaskID() uuid.UUID
	Type() string
	Created(CreatedBy *uuid.UUID, CreatedAt *time.Time)
}

type TaskEventImpl struct {
	event.DomainEvent
	taskID    *uuid.UUID
	eventType string
	CreatedBy *uuid.UUID
	CreatedAt *time.Time
}

func newTaskEvent(taskID *uuid.UUID, eventType string) TaskEvent {
	return &TaskEventImpl{
		DomainEvent: event.NewDomainEvent(),
		taskID:      taskID,
		eventType:   eventType,
	}
}

func (t *TaskEventImpl) TaskID() uuid.UUID {
	return *t.taskID
}

func (t *TaskEventImpl) Type() string {
	return t.eventType
}

func (t *TaskEventImpl) Created(CreatedBy *uuid.UUID, CreatedAt *time.Time) {
	t.CreatedBy = CreatedBy
	t.CreatedAt = CreatedAt
}
