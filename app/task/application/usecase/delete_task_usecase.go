package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
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
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(ctx, "指定されたタスクが見つかりませんでした。")
	}

	if err := u.taskRepository.Delete(ctx, task); err != nil {
		return xerrors.Errorf("taskRepository.Delete(): %w", err)
	}

	if err := u.taskPresenter.NoContentResponse(ctx); err != nil {
		return xerrors.Errorf("taskPresenter.NoContentResponse(): %w", err)
	}

	a := auth.GetAuth(ctx)
	taskDeleted := event.NewTaskDeleted(task, &a.UserId)
	err = u.taskEventRepository.Register(ctx, taskDeleted)
	if err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	u.publisher.Publish(taskDeleted)

	return nil
}
