package event

import (
	"github.com/google/uuid"
)

type TaskUpdated struct {
	*TaskEvent
}

func NewTaskUpdated(taskID *uuid.UUID) *TaskUpdated {
	return &TaskUpdated{TaskEvent: NewTaskEvent(taskID)}
}
