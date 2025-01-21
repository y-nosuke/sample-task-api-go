package event

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
)

type TaskEdited struct {
	TaskEvent
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
}

func NewTaskEdited(taskID uuid.UUID, title string, detail *string, completed bool, deadline *time.Time, editedBy uuid.UUID, editedAt time.Time) (*TaskEdited, error) {
	taskEvent, err := newTaskEvent(taskID, editedBy, editedAt)
	if err != nil {
		return nil, xerrors.Errorf("newTaskEvent(): %w", err)
	}
	return &TaskEdited{
		TaskEvent: *taskEvent,
		Title:     title,
		Detail:    detail,
		Completed: completed,
		Deadline:  deadline,
	}, nil
}
