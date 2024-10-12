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
	fmt.Println("タスク完了処理を開始します。")

	task, err := u.taskRepository.GetById(cctx, args.Id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			if resErr := u.taskPresenter.NotFound(cctx, "指定されたタスクが見つかりませんでした。"); resErr != nil {
				return xerrors.Errorf("original error: %v, taskPresenter.NotFound(): %w", err, resErr)
			}
			return ferrors.NewBusinessErrorf(err, "指定されたタスクが見つかりませんでした。")
		}
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	taskCompleted := task.Complete(auth.GetUserId(cctx))

	if err = u.taskRepository.Update(cctx, task, args.Version); err != nil {
		if errors.Is(err, repository.ErrNotAffected) {
			if resErr := u.taskPresenter.Conflict(cctx, "タスクは既に更新済みです。"); resErr != nil {
				return xerrors.Errorf("original error: %v, taskPresenter.Conflict(): %w", err, resErr)
			}
			return ferrors.NewBusinessErrorf(err, "タスクは既に更新済みです。")
		}
		return xerrors.Errorf("taskRepository.Update(): %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 task: %+v\n", task)

	if err = u.taskEventRepository.RegisterTaskCompleted(cctx, taskCompleted); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskCompleted: %+v\n", taskCompleted)

	u.publisher.Publish(taskCompleted)

	if err = u.taskPresenter.NilResponse(cctx); err != nil {
		return xerrors.Errorf("taskPresenter.NilResponse(): %w", err)
	}

	return nil
}
