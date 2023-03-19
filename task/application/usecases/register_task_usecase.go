package usecases

import (
	"context"
	"time"

	"github.com/y-nosuke/sample-task-api-go/task/application/presenters"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
	"golang.org/x/xerrors"
)

type RegisterTaskUseCaseArgs struct {
	Title    string
	Detail   *string
	Deadline *time.Time
}

type RegisterTaskUseCase struct {
	taskRepository repositories.TaskRepository
	taskPresenter  presenters.TaskPresenter
}

func NewRegisterTaskUseCase(taskRepository repositories.TaskRepository, taskPresenter presenters.TaskPresenter) *RegisterTaskUseCase {
	return &RegisterTaskUseCase{taskRepository, taskPresenter}
}

func (u *RegisterTaskUseCase) Invoke(ctx context.Context, args *RegisterTaskUseCaseArgs) error {
	task := entities.NewTask(args.Title, args.Detail, args.Deadline)

	if err := u.taskRepository.Register(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := u.taskPresenter.RegisterTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
