package mapping

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

func RTask(task *entity.Task, userId *uuid.UUID, version *uuid.UUID) (*dao.RTask, error) {
	id, err := task.Id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("task.Id.MarshalBinary(): %w", err)
	}

	var byteVersion []byte
	if version != nil {
		byteVersion, err = version.MarshalBinary()
		if err != nil {
			return nil, xerrors.Errorf("version.MarshalBinary(): %w", err)
		}
	}

	byteUserId, err := userId.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("userId.MarshalBinary(): %w", err)
	}

	return &dao.RTask{
		ID:        id,
		Title:     task.Title,
		Detail:    null.StringFromPtr(task.Detail),
		Completed: task.Completed,
		Deadline:  null.TimeFromPtr(task.Deadline),
		CreatedBy: byteUserId,
		UpdatedBy: byteUserId,
		Version:   byteVersion,
	}, nil
}

func TaskSlice(rTaskSlice dao.RTaskSlice) (entity.TaskSlice, error) {
	var taskSlice entity.TaskSlice
	for _, t := range rTaskSlice {
		task, err := Task(t)
		if err != nil {
			return nil, xerrors.Errorf("Task(): %w", err)
		}

		taskSlice = append(taskSlice, task)
	}

	return taskSlice, nil
}

func Task(rTask *dao.RTask) (*entity.Task, error) {
	id, err := uuid.FromBytes(rTask.ID)
	if err != nil {
		return nil, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}

	var detail *string
	if rTask.Detail.Valid {
		detail = &rTask.Detail.String
	}

	var deadline *time.Time
	if rTask.Deadline.Valid {
		deadline = &rTask.Deadline.Time
	}

	createdBy, err := uuid.FromBytes(rTask.CreatedBy)
	if err != nil {
		return nil, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}

	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return nil, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}

	version, err := uuid.FromBytes(rTask.Version)
	if err != nil {
		return nil, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}

	return &entity.Task{
		Id:        &id,
		Title:     rTask.Title,
		Detail:    detail,
		Completed: rTask.Completed,
		Deadline:  deadline,
		CreatedBy: &createdBy,
		CreatedAt: &rTask.CreatedAt,
		UpdatedBy: &updatedBy,
		UpdatedAt: &rTask.UpdatedAt,
		Version:   &version,
	}, nil
}

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
