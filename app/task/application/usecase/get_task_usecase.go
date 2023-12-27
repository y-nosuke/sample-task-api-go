package usecase

import (
	"context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
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
	if err != nil {
		return errors.SystemErrorf("taskRepository.GetById(): %w", err)
	}

	if task == nil {
		if err := u.taskPresenter.NotFound(ctx, "指定されたタスクが見つかりませんでした。"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	if err := u.taskPresenter.GetTaskResponse(ctx, task); err != nil {
		return errors.SystemErrorf("taskPresenter.GetTaskResponse(): %w", err)
	}

	return nil
}
