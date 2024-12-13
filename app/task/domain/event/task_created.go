package event

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
)

type TaskCreated struct {
	TaskEvent
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
}

func NewTaskCreated(taskID uuid.UUID, title string, detail *string, completed bool, deadline *time.Time, createdBy uuid.UUID, createdAt time.Time) (*TaskCreated, error) {
	taskEvent, err := newTaskEvent(taskID, createdBy, createdAt)
	if err != nil {
		return nil, xerrors.Errorf("newTaskEvent(): %w", err)
	}
	return &TaskCreated{
		TaskEvent: *taskEvent,
		Title:     title,
		Detail:    detail,
		Completed: completed,
		Deadline:  deadline,
	}, nil
}
