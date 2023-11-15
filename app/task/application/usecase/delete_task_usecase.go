package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type DeleteTaskUseCaseArgs struct {
	Id uuid.UUID
}

type DeleteTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewDeleteTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{taskRepository, taskPresenter}
}

func (u *DeleteTaskUseCase) Invoke(ctx context.Context, args *DeleteTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(ctx, "指定されたタスクが見つかりませんでした。")
	}

	if err := u.taskRepository.Delete(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.NoContentResponse(ctx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
