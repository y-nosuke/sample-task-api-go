package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

func RTask(task *entity.Task, userId uuid.UUID, version uuid.UUID, create bool) (*dao.RTask, error) {
	id, err := task.Id().MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("task.id.MarshalBinary(): %w", err)
	}

	byteCreatedBy, err := task.CreatedBy().MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("userId.MarshalBinary(): %w", err)
	}

	byteEditedBy, err := task.EditedBy().MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("userId.MarshalBinary(): %w", err)
	}

	byteUserId, err := userId.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("userId.MarshalBinary(): %w", err)
	}

	byteVersion, err := version.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("version.MarshalBinary(): %w", err)
	}

	rTask := &dao.RTask{
		ID:        id,
		Title:     task.Title(),
		Detail:    null.StringFromPtr(task.Detail()),
		Completed: task.Completed(),
		Deadline:  null.TimeFromPtr(task.Deadline()),
		CreatedBy: byteCreatedBy,
		CreatedAt: task.CreatedAt(),
		EditedBy:  byteEditedBy,
		EditedAt:  task.EditedAt(),
		UpdatedBy: byteUserId,
		Version:   byteVersion,
	}

	if create {
		rTask.RegisterBy = byteUserId
	}

	return rTask, nil
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

	return entity.NewTask(id, rTask.Title, detail, rTask.Completed, deadline, createdBy, rTask.CreatedAt, updatedBy, rTask.UpdatedAt, version), nil
}
