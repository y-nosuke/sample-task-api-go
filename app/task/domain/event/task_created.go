package event

import (
	"time"

	"github.com/google/uuid"
)

type TaskCreated struct {
	TaskEvent
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
}

func NewTaskCreated(taskID uuid.UUID, title string, detail *string, completed bool, deadline *time.Time, createdBy uuid.UUID, createdAt time.Time) *TaskCreated {
	return &TaskCreated{
		TaskEvent: *newTaskEvent(taskID, createdBy, createdAt),
		Title:     title,
		Detail:    detail,
		Completed: completed,
		Deadline:  deadline,
	}
}
