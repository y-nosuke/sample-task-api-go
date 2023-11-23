package event

import (
	"github.com/google/uuid"
)

type TaskUnCompleted struct {
	*TaskEvent
}

func NewTaskUnCompleted(taskID *uuid.UUID) *TaskUnCompleted {
	return &TaskUnCompleted{TaskEvent: NewTaskEvent(taskID)}
}
