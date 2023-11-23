package event

import (
	"github.com/google/uuid"
)

type TaskCompleted struct {
	*TaskEvent
}

func NewTaskCompleted(taskID *uuid.UUID) *TaskCompleted {
	return &TaskCompleted{TaskEvent: NewTaskEvent(taskID)}
}
