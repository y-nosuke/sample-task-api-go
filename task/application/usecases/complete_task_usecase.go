package usecases

import (
	"context"
	"github.com/google/uuid"
	ferrors "github.com/y-nosuke/sample-task-api-go/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/task/application/presenters"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"golang.org/x/xerrors"
)

type CompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version *uuid.UUID
}

type CompleteTaskUseCase struct {
	taskRepository repositories.TaskRepository
	taskPresenter  presenters.TaskPresenter
}

func NewCompleteTaskUseCase(taskRepository repositories.TaskRepository, taskPresenter presenters.TaskPresenter) *CompleteTaskUseCase {
	return &CompleteTaskUseCase{taskRepository, taskPresenter}
}

func (u *CompleteTaskUseCase) Invoke(ctx context.Context, args *CompleteTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if task == nil {
		return ferrors.New(ferrors.NotFound, "指定されたタスクが見つかりませんでした。", err)
	} else if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	task.Complete(args.Version)

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
