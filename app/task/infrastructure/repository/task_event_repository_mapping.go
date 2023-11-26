package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

func ETaskEvent(taskEvent event.TaskEvent, userId *uuid.UUID) (*dao.ETaskEvent, error) {
	id, err := taskEvent.ID().MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("task.ID().MarshalBinary(): %w", err)
	}

	eventType, err := taskEventType(taskEvent)
	if err != nil {
		return nil, xerrors.Errorf("taskEventType(): %w", err)
	}

	taskID, err := taskEvent.TaskID().MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("task.TaskID().MarshalBinary(): %w", err)
	}

	data, err := json.Marshal(taskEvent.Data())
	if err != nil {
		return nil, xerrors.Errorf("json.Marshal(): %w", err)
	}

	byteUserId, err := userId.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("userId.MarshalBinary(): %w", err)
	}

	return &dao.ETaskEvent{
		ID:        id,
		Type:      eventType,
		TaskID:    taskID,
		Data:      data,
		CreatedBy: byteUserId,
	}, nil
}

const (
	ETaskCreated     = "TaskCreated"
	ETaskUpdated     = "TaskUpdated"
	ETaskDeleted     = "TaskDeleted"
	ETaskCompleted   = "TaskCompleted"
	ETaskUnCompleted = "TaskUnCompleted"
)

func taskEventType(taskEvent event.TaskEvent) (string, error) {
	if _, ok := taskEvent.(*event.TaskCreated); ok {
		return ETaskCreated, nil
	} else if _, ok := taskEvent.(*event.TaskUpdated); ok {
		return ETaskUpdated, nil
	} else if _, ok := taskEvent.(*event.TaskCompleted); ok {
		return ETaskCompleted, nil
	} else if _, ok := taskEvent.(*event.TaskUnCompleted); ok {
		return ETaskUnCompleted, nil
	} else if _, ok := taskEvent.(*event.TaskDeleted); ok {
		return ETaskDeleted, nil
	} else {
		return "", xerrors.Errorf("unknown event. eventID: %s\n", taskEvent.ID())
	}
}
