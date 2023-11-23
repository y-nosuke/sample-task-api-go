package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
)

type TaskEvent struct {
	event.DomainEvent
	taskID *uuid.UUID
}

func NewTaskEvent(taskID *uuid.UUID) *TaskEvent {
	return &TaskEvent{DomainEvent: event.NewDomainEvent(), taskID: taskID}
}

func (t TaskEvent) TaskID() uuid.UUID {
	return *t.taskID
}
