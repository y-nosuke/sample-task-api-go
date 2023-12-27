package usecase

import (
	"context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
)

type GetAllTaskUseCaseArgs struct {
}

type GetAllTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewGetAllTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *GetAllTaskUseCase {
	return &GetAllTaskUseCase{taskRepository, taskPresenter}
}

func (u *GetAllTaskUseCase) Invoke(ctx context.Context, _ *GetAllTaskUseCaseArgs) error {
	tasks, err := u.taskRepository.GetAll(ctx)
	if err != nil {
		return errors.SystemErrorf("taskRepository.GetAll(): %w", err)
	}

	if err := u.taskPresenter.TaskAllResponse(ctx, tasks); err != nil {
		return errors.SystemErrorf("taskPresenter.TaskAllResponse(): %w", err)
	}

	return nil
}
