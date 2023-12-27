package usecase

import (
	"context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/google/uuid"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
)

type UnCompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version *uuid.UUID
}

type UnCompleteTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewUnCompleteTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *UnCompleteTaskUseCase {
	return &UnCompleteTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *UnCompleteTaskUseCase) Invoke(ctx context.Context, args *UnCompleteTaskUseCaseArgs) error {
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

	task.UnComplete(args.Version)

	if row, err := u.taskRepository.Update(ctx, task, args.Version); err != nil {
		return errors.SystemErrorf("taskRepository.Update(): %w", err)
	} else if row != 1 {
		if err := u.taskPresenter.Conflict(ctx, "タスクは既に更新済みです。"); err != nil {
			return errors.SystemErrorf("taskPresenter.Conflict(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Conflict()")
	}

	if err := u.taskPresenter.NilResponse(ctx); err != nil {
		return errors.SystemErrorf("taskPresenter.NilResponse(): %w", err)
	}

	taskUnCompleted := event.NewTaskUnCompleted(task)
	err = u.taskEventRepository.Register(ctx, taskUnCompleted)
	if err != nil {
		return errors.SystemErrorf("taskEventRepository.Register(): %w", err)
	}
	u.publisher.Publish(taskUnCompleted)

	return nil
}
