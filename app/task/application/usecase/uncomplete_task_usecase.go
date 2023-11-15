package usecase

import (
	"context"
	"github.com/google/uuid"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
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
	if task == nil {
		return ferrors.New(ferrors.NotFound, "指定されたタスクが見つかりませんでした。", err)
	} else if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	task.UnComplete(args.Version)

	if row, err := u.taskRepository.Update(ctx, task); row != 1 {
		return ferrors.New(ferrors.Conflict, "タスクは既に更新済みです。", err)
	} else if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.NilResponse(ctx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
