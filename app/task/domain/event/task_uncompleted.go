package event

import (
	"time"

	"github.com/google/uuid"
)

type TaskUnCompleted struct {
	TaskEvent
}

func NewTaskUnCompleted(taskID uuid.UUID, editedBy uuid.UUID, editedAt time.Time) *TaskUnCompleted {
	return &TaskUnCompleted{
		TaskEvent: *newTaskEvent(taskID, editedBy, editedAt),
	}
}
