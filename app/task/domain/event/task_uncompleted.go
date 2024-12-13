package event

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
)

type TaskUnCompleted struct {
	TaskEvent
}

func NewTaskUnCompleted(taskID uuid.UUID, editedBy uuid.UUID, editedAt time.Time) (*TaskUnCompleted, error) {
	taskEvent, err := newTaskEvent(taskID, editedBy, editedAt)
	if err != nil {
		return nil, xerrors.Errorf("newTaskEvent(): %w", err)
	}
	return &TaskUnCompleted{
		TaskEvent: *taskEvent,
	}, nil
}
