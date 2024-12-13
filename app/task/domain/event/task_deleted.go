package event

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
)

type TaskDeleted struct {
	TaskEvent
}

func NewTaskDeleted(taskID uuid.UUID, deletedBy uuid.UUID) (*TaskDeleted, error) {
	now := time.Now()
	taskEvent, err := newTaskEvent(taskID, deletedBy, now)
	if err != nil {
		return nil, xerrors.Errorf("newTaskEvent(): %w", err)
	}
	return &TaskDeleted{
		TaskEvent: *taskEvent,
	}, nil
}
