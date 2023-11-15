package usecase

import (
	"context"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"time"

	"golang.org/x/xerrors"
)

type RegisterTaskUseCaseArgs struct {
	Title    string
	Detail   *string
	Deadline *time.Time
}

type RegisterTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewRegisterTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *RegisterTaskUseCase {
	return &RegisterTaskUseCase{taskRepository, taskPresenter}
}

func (u *RegisterTaskUseCase) Invoke(ctx context.Context, args *RegisterTaskUseCaseArgs) error {
	task := entity.NewTask(args.Title, args.Detail, args.Deadline)

	if err := u.taskRepository.Register(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.RegisterTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
