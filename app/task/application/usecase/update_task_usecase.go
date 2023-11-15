package usecase

import (
	"context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UpdateTaskUseCaseArgs struct {
	Id       uuid.UUID
	Title    string
	Detail   *string
	Deadline *time.Time
	Version  *uuid.UUID
}

type UpdateTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewUpdateTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{taskRepository, taskPresenter}
}

func (u *UpdateTaskUseCase) Invoke(ctx context.Context, args *UpdateTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(ctx, args.Id)
	if task == nil {
		return ferrors.New(ferrors.NotFound, "指定されたタスクが見つかりませんでした。", err)
	} else if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	task.Update(args.Title, args.Detail, args.Deadline, args.Version)

	if row, err := u.taskRepository.Update(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if row != 1 {
		return ferrors.New(ferrors.Conflict, "タスクは既に更新済みです。", err)
	}

	if err := u.taskPresenter.UpdateTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
