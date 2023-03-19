package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/task/application/presenters"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"golang.org/x/xerrors"
)

type DeleteTaskUseCaseArgs struct {
	Id uuid.UUID
}

type DeleteTaskUseCase struct {
	taskRepository repositories.TaskRepository
	taskPresenter  presenters.TaskPresenter
}

func NewDeleteTaskUseCase(taskRepository repositories.TaskRepository, taskPresenter presenters.TaskPresenter) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{taskRepository, taskPresenter}
}

func (u *DeleteTaskUseCase) Invoke(ctx context.Context, args *DeleteTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskRepository.Delete(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.NilResponse(ctx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
