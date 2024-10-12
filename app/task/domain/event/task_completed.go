package event

import (
	"time"

	"github.com/google/uuid"
)

type TaskCompleted struct {
	TaskEvent
}

func NewTaskCompleted(taskID uuid.UUID, editedBy uuid.UUID, editedAt time.Time) *TaskCompleted {
	return &TaskCompleted{
		TaskEvent: *newTaskEvent(taskID, editedBy, editedAt),
	}
}
