package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
)

type TaskEvent interface {
	event.DomainEvent
	TaskID() uuid.UUID
	Data() any
	Created(CreatedBy uuid.UUID, CreatedAt time.Time)
}

type TaskEventCommon struct {
	event.DomainEvent
	taskID    uuid.UUID
	CreatedBy uuid.UUID
	CreatedAt time.Time
}

func newTaskEventCommon(taskID uuid.UUID) *TaskEventCommon {
	return &TaskEventCommon{
		DomainEvent: event.NewDomainEvent(),
		taskID:      taskID,
	}
}

func (t *TaskEventCommon) TaskID() uuid.UUID {
	return t.taskID
}

func (t *TaskEventCommon) Created(CreatedBy uuid.UUID, CreatedAt time.Time) {
	t.CreatedBy = CreatedBy
	t.CreatedAt = CreatedAt
}
