package event

import (
	"time"

	"github.com/google/uuid"
)

type TaskDeleted struct {
	TaskEvent
}

func NewTaskDeleted(taskID uuid.UUID, deletedBy uuid.UUID) *TaskDeleted {
	now := time.Now()
	return &TaskDeleted{
		TaskEvent: *newTaskEvent(taskID, deletedBy, now),
	}
}
