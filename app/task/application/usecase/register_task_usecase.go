package usecase

import (
	"context"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
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
	publisher      observer.Publisher[nevent.DomainEvent]
}

func NewRegisterTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *RegisterTaskUseCase {
	return &RegisterTaskUseCase{taskRepository, taskPresenter, publisher}
}

func (u *RegisterTaskUseCase) Invoke(ctx context.Context, args *RegisterTaskUseCaseArgs) error {
	task := entity.NewTask(args.Title, args.Detail, args.Deadline)

	if err := u.taskRepository.Register(ctx, task); err != nil {
		return xerrors.Errorf("taskRepository.Register(): %w", err)
	}

	if err := u.taskPresenter.RegisterTaskResponse(ctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.RegisterTaskResponse(): %w", err)
	}

	taskCreated := event.NewTaskCreated(task.Id)
	u.publisher.Publish(taskCreated)

	return nil
}
