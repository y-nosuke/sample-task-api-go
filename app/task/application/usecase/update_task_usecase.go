package usecase

import (
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
	"time"
)

type UpdateTaskUseCaseArgs struct {
	Id       uuid.UUID
	Title    string
	Detail   *string
	Deadline *time.Time
	Version  uuid.UUID
}

type UpdateTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewUpdateTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *UpdateTaskUseCase) Invoke(cctx fcontext.Context, args *UpdateTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(cctx, args.Id)
	if err != nil {
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(cctx, "指定されたタスクが見つかりませんでした。")
	}

	task.Update(args.Title, args.Detail, args.Deadline, args.Version)

	if row, err := u.taskRepository.Update(cctx, task, args.Version); err != nil {
		return xerrors.Errorf("taskRepository.Update(): %w", err)
	} else if row != 1 {
		return u.taskPresenter.Conflict(cctx, "タスクは既に更新済みです。")
	}

	if err := u.taskPresenter.UpdateTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.UpdateTaskResponse(): %w", err)
	}

	taskUpdated := event.NewTaskUpdated(task)
	err = u.taskEventRepository.Register(cctx, taskUpdated)
	if err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	u.publisher.Publish(taskUpdated)

	return nil
}
