package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type CompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version *uuid.UUID
}

type CompleteTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewCompleteTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *CompleteTaskUseCase {
	return &CompleteTaskUseCase{taskRepository, taskPresenter}
}

func (u *CompleteTaskUseCase) Invoke(ctx context.Context, args *CompleteTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(ctx, "指定されたタスクが見つかりませんでした。")
	}

	task.Complete(args.Version)

	// TODO 重複エラーは独自errorを返すようにする
	if row, err := u.taskRepository.Update(ctx, task, args.Version); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if row != 1 {
		return u.taskPresenter.Conflict(ctx, "タスクは既に更新済みです。")
	}

	if err := u.taskPresenter.NilResponse(ctx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
