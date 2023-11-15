package mapping

import (
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
	"time"
)

func TaskDto(task *entity.Task, userId *uuid.UUID) (*dao.Task, error) {
	id, err := task.Id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var version []byte
	if task.Version != nil {
		version, err = task.Version.MarshalBinary()
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
	}

	uid, err := userId.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &dao.Task{
		ID:        id,
		Title:     task.Title,
		Detail:    null.StringFromPtr(task.Detail),
		Completed: task.Completed,
		Deadline:  null.TimeFromPtr(task.Deadline),
		CreatedBy: uid,
		UpdatedBy: uid,
		Version:   version,
	}, nil
}

func TaskSlice(taskSlice dao.TaskSlice) (entity.TaskSlice, error) {
	var tasks entity.TaskSlice
	for _, t := range taskSlice {
		task, err := Task(t)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Task(taskDto *dao.Task) (*entity.Task, error) {
	id, err := uuid.FromBytes(taskDto.ID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var detail *string
	if taskDto.Detail.Valid {
		detail = &taskDto.Detail.String
	}

	var deadline *time.Time
	if taskDto.Deadline.Valid {
		deadline = &taskDto.Deadline.Time
	}

	createdBy, err := uuid.FromBytes(taskDto.CreatedBy)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	updatedBy, err := uuid.FromBytes(taskDto.UpdatedBy)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &entity.Task{
		Id:        &id,
		Title:     taskDto.Title,
		Detail:    detail,
		Completed: taskDto.Completed,
		Deadline:  deadline,
		CreatedBy: &createdBy,
		CreatedAt: &taskDto.CreatedAt,
		UpdatedBy: &updatedBy,
		UpdatedAt: &taskDto.UpdatedAt,
		Version:   &version,
	}, nil
}
