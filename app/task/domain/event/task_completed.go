package event

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
)

type TaskCompleted struct {
	TaskEvent
}

func NewTaskCompleted(taskID uuid.UUID, editedBy uuid.UUID, editedAt time.Time) (*TaskCompleted, error) {
	taskEvent, err := newTaskEvent(taskID, editedBy, editedAt)
	if err != nil {
		return nil, xerrors.Errorf("newTaskEvent(): %w", err)
	}
	return &TaskCompleted{
		TaskEvent: *taskEvent,
	}, nil
}
