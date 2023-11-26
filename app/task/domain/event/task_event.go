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
	Data() any
	Created(CreatedBy *uuid.UUID, CreatedAt *time.Time)
}

type TaskEventCommon struct {
	event.DomainEvent
	taskID    *uuid.UUID
	eventType string
	CreatedBy *uuid.UUID
	CreatedAt *time.Time
}

func newTaskEventCommon(taskID *uuid.UUID, eventType string) *TaskEventCommon {
	return &TaskEventCommon{
		DomainEvent: event.NewDomainEvent(),
		taskID:      taskID,
		eventType:   eventType,
	}
}

func (t *TaskEventCommon) TaskID() uuid.UUID {
	return *t.taskID
}

func (t *TaskEventCommon) Type() string {
	return t.eventType
}

func (t *TaskEventCommon) Created(CreatedBy *uuid.UUID, CreatedAt *time.Time) {
	t.CreatedBy = CreatedBy
	t.CreatedAt = CreatedAt
}
