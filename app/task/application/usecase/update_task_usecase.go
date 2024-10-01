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
		if err = u.taskPresenter.NotFound(cctx, "指定されたタスクが見つかりませんでした。"); err != nil {
			return xerrors.Errorf("taskPresenter.NotFound(): %w", err)
		}
		return nil
	}

	task.Update(args.Title, args.Detail, args.Deadline, args.Version)

	var row int
	if row, err = u.taskRepository.Update(cctx, task, args.Version); err != nil {
		return xerrors.Errorf("taskRepository.Update(): %w", err)
	} else if row != 1 {
		if err = u.taskPresenter.Conflict(cctx, "タスクは既に更新済みです。"); err != nil {
			return err
		}
		return nil
	}

	taskUpdated := event.NewTaskUpdated(task)
	if err = u.taskEventRepository.Register(cctx, taskUpdated); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	u.publisher.Publish(taskUpdated)

	if err = u.taskPresenter.UpdateTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.UpdateTaskResponse(): %w", err)
	}

	return nil
}
