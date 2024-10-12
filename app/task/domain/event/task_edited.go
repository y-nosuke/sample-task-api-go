package event

import (
	"time"

	"github.com/google/uuid"
)

type TaskEdited struct {
	TaskEvent
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
}

func NewTaskUpdated(taskID uuid.UUID, title string, detail *string, completed bool, deadline *time.Time, editedBy uuid.UUID, editedAt time.Time) *TaskEdited {
	return &TaskEdited{
		TaskEvent: *newTaskEvent(taskID, editedBy, editedAt),
		Title:     title,
		Detail:    detail,
		Completed: completed,
		Deadline:  deadline,
	}
}
