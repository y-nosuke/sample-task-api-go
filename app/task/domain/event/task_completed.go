package event

import (
	"github.com/google/uuid"
	"time"
)

type TaskCompleted struct {
	TaskEvent
}

func NewTaskCompleted(taskID uuid.UUID, editedBy uuid.UUID, editedAt time.Time) *TaskCompleted {
	return &TaskCompleted{
		TaskEvent: *newTaskEvent(taskID, editedBy, editedAt),
	}
}
