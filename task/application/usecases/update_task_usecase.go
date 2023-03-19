package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/task/application/presenters"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"golang.org/x/xerrors"
)

type UpdateTaskUseCaseArgs struct {
	Id        uuid.UUID
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
	Version   *uuid.UUID
}

type UpdateTaskUseCase struct {
	taskRepository repositories.TaskRepository
	taskPresenter  presenters.TaskPresenter
}

func NewUpdateTaskUseCase(taskRepository repositories.TaskRepository, taskPresenter presenters.TaskPresenter) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{taskRepository, taskPresenter}
}

func (u *UpdateTaskUseCase) Invoke(ctx context.Context, args *UpdateTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	task.Update(args.Title, args.Detail, args.Completed, args.Deadline, args.Version)

	if err := u.taskRepository.Update(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.UpdateTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
