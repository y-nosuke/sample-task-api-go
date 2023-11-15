package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type UnCompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version *uuid.UUID
}

type UnCompleteTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewUnCompleteTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *UnCompleteTaskUseCase {
	return &UnCompleteTaskUseCase{taskRepository, taskPresenter}
}

func (u *UnCompleteTaskUseCase) Invoke(ctx context.Context, args *UnCompleteTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(ctx, "指定されたタスクが見つかりませんでした。")
	}

	task.UnComplete(args.Version)

	if row, err := u.taskRepository.Update(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if row != 1 {
		return u.taskPresenter.Conflict(ctx, "タスクは既に更新済みです。")
	}

	if err := u.taskPresenter.NilResponse(ctx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}