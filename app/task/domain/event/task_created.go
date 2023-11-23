package event

import (
	"github.com/google/uuid"
)

type TaskCreated struct {
	*TaskEvent
}

func NewTaskCreated(taskID *uuid.UUID) *TaskCreated {
	return &TaskCreated{TaskEvent: NewTaskEvent(taskID)}
}
