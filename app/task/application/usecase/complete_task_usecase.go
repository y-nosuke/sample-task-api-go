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

type CompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version *uuid.UUID
}

type CompleteTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewCompleteTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *CompleteTaskUseCase {
	return &CompleteTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *CompleteTaskUseCase) Invoke(ctx context.Context, args *CompleteTaskUseCaseArgs) error {
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

	task.Complete(args.Version)

	// TODO 重複エラーは独自errorを返すようにする
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

	taskCompleted := event.NewTaskCompleted(task)
	err = u.taskEventRepository.Register(ctx, taskCompleted)
	if err != nil {
		return errors.SystemErrorf("taskEventRepository.Register(): %w", err)
	}
	u.publisher.Publish(taskCompleted)

	return nil
}
