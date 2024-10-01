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
)

type CompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version uuid.UUID
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

func (u *CompleteTaskUseCase) Invoke(cctx fcontext.Context, args *CompleteTaskUseCaseArgs) error {
	task, err := u.taskRepository.GetById(cctx, args.Id)
	if err != nil {
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	if task == nil {
		return u.taskPresenter.NotFound(cctx, "指定されたタスクが見つかりませんでした。")
	}

	task.Complete(args.Version)

	// TODO 重複エラーは独自errorを返すようにする
	if row, err := u.taskRepository.Update(cctx, task, args.Version); err != nil {
		return xerrors.Errorf("taskRepository.Update(): %w", err)
	} else if row != 1 {
		return u.taskPresenter.Conflict(cctx, "タスクは既に更新済みです。")
	}

	if err := u.taskPresenter.NilResponse(cctx); err != nil {
		return xerrors.Errorf("taskPresenter.NilResponse(): %w", err)
	}

	taskCompleted := event.NewTaskCompleted(task)
	err = u.taskEventRepository.Register(cctx, taskCompleted)
	if err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}
	u.publisher.Publish(taskCompleted)

	return nil
}
