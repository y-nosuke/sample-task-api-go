package usecase

import (
	"context"
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
	if err != nil {
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(ctx, "指定されたタスクが見つかりませんでした。")
	}

	task.Update(args.Title, args.Detail, args.Deadline, args.Version)

	if row, err := u.taskRepository.Update(ctx, task, args.Version); err != nil {
		return xerrors.Errorf("taskRepository.Update(): %w", err)
	} else if row != 1 {
		return u.taskPresenter.Conflict(ctx, "タスクは既に更新済みです。")
	}

	if err := u.taskPresenter.UpdateTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.UpdateTaskResponse(): %w", err)
	}

	return nil
}
