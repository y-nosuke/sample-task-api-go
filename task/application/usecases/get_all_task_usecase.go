package usecases

import (
	"context"

	"github.com/y-nosuke/sample-task-api-go/task/application/presenters"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"golang.org/x/xerrors"
)

type GetAllTaskUseCaseArgs struct {
}

type GetAllTaskUseCase struct {
	taskRepository repositories.TaskRepository
	taskPresenter  presenters.TaskPresenter
}

func NewGetAllTaskUseCase(taskRepository repositories.TaskRepository, taskPresenter presenters.TaskPresenter) *GetAllTaskUseCase {
	return &GetAllTaskUseCase{taskRepository, taskPresenter}
}

func (u *GetAllTaskUseCase) Invoke(ctx context.Context, args *GetAllTaskUseCaseArgs) error {
	tasks, err := u.taskRepository.GetAll(ctx)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.TaskAllResponse(ctx, tasks); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
