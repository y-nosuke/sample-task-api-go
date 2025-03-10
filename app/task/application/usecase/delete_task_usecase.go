package usecase

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
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

func (u *DeleteTaskUseCase) Invoke(cctx fcontext.Context, args *DeleteTaskUseCaseArgs) error {
	fmt.Println("タスク削除処理を開始します。")

	task, err := u.taskRepository.GetById(cctx, args.Id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ferrors.NewNotFoundErrorf(err, "指定されたタスクが見つかりませんでした。")
		}
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	if err = u.taskRepository.Delete(cctx, task); err != nil {
		if errors.Is(err, repository.ErrNotAffected) {
			return ferrors.NewConflictErrorf(err, "タスクは既に更新済みです。")
		}
		return xerrors.Errorf("taskRepository.Delete(): %w", err)
	}

	fmt.Printf("データベースのタスクが削除されました。 task: %+v\n", task)

	taskDeleted, err := event.NewTaskDeleted(task.Id(), auth.GetUserId(cctx))
	if err != nil {
		return xerrors.Errorf("event.NewTaskDeleted(): %w", err)
	}

	if err = u.taskEventRepository.RegisterTaskDeleted(cctx, taskDeleted); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskDeleted: %+v\n", taskDeleted)

	u.publisher.Publish(taskDeleted)

	if err = u.taskPresenter.NoContentResponse(cctx); err != nil {
		return xerrors.Errorf("taskPresenter.NoContentResponse(): %w", err)
	}

	return nil
}
