package event

import (
	"github.com/google/uuid"
)

type TaskDeleted struct {
	*TaskEvent
}

func NewTaskDeleted(taskID *uuid.UUID) *TaskDeleted {
	return &TaskDeleted{TaskEvent: NewTaskEvent(taskID)}
}
