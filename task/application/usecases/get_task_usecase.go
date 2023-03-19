package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/task/application/presenters"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"golang.org/x/xerrors"
)

type GetTaskUseCaseArgs struct {
	Id uuid.UUID
}

type GetTaskUseCase struct {
	taskRepository repositories.TaskRepository
	taskPresenter  presenters.TaskPresenter
}

func NewGetTaskUseCase(taskRepository repositories.TaskRepository, taskPresenter presenters.TaskPresenter) *GetTaskUseCase {
	return &GetTaskUseCase{taskRepository, taskPresenter}
}

func (u *GetTaskUseCase) Invoke(ctx context.Context, args *GetTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.GetTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
