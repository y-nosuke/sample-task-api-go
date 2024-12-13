package factory

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

func CreateTask(title string, detail *string, deadline *time.Time, userID uuid.UUID) (*entity.Task, *event.TaskCreated, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, nil, xerrors.Errorf("uuid.NewV7(): %w", err)
	}
	now := time.Now()
	task := entity.NewTask(id, title, detail, false, deadline, userID, now, userID, now, uuid.Nil)
	created, err := taskCreated(task)
	if err != nil {
		return nil, nil, xerrors.Errorf("taskCreated(): %w", err)
	}

	return task, created, nil
}

func taskCreated(task *entity.Task) (*event.TaskCreated, error) {
	created, err := event.NewTaskCreated(task.Id(), task.Title(), task.Detail(), task.Completed(), task.Deadline(), task.CreatedBy(), task.CreatedAt())
	if err != nil {
		return nil, xerrors.Errorf("event.NewTaskCreated(): %w", err)
	}
	return created, nil
}
