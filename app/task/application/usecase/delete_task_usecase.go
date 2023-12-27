package usecase

import (
	"context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
)

type DeleteTaskUseCaseArgs struct {
	Id uuid.UUID
}

type DeleteTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewDeleteTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *DeleteTaskUseCase) Invoke(ctx context.Context, args *DeleteTaskUseCaseArgs) error {
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

	if err := u.taskRepository.Delete(ctx, task); err != nil {
		return errors.SystemErrorf("taskRepository.Delete(): %w", err)
	}

	if err := u.taskPresenter.NoContentResponse(ctx); err != nil {
		return errors.SystemErrorf("taskPresenter.NoContentResponse(): %w", err)
	}

	a := auth.GetAuth(ctx)
	taskDeleted := event.NewTaskDeleted(task, &a.UserId)
	err = u.taskEventRepository.Register(ctx, taskDeleted)
	if err != nil {
		return errors.SystemErrorf("taskEventRepository.Register(): %w", err)
	}

	u.publisher.Publish(taskDeleted)

	return nil
}
