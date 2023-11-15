package usecase

import (
	"context"
	"github.com/google/uuid"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type GetTaskUseCaseArgs struct {
	Id uuid.UUID
}

type GetTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewGetTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *GetTaskUseCase {
	return &GetTaskUseCase{taskRepository, taskPresenter}
}

func (u *GetTaskUseCase) Invoke(ctx context.Context, args *GetTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if task == nil {
		return ferrors.New(ferrors.NotFound, "指定されたタスクが見つかりませんでした。", err)
	} else if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.GetTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
