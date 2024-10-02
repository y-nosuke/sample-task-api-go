package usecase

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type UnCompleteTaskUseCaseArgs struct {
	Id      uuid.UUID
	Version uuid.UUID
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

func (u *UnCompleteTaskUseCase) Invoke(cctx fcontext.Context, args *UnCompleteTaskUseCaseArgs) error {
	fmt.Println("タスク未完了処理を開始します。")

	task, err := u.taskRepository.GetById(cctx, args.Id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			if err = u.taskPresenter.NotFound(cctx, "指定されたタスクが見つかりませんでした。"); err != nil {
				return xerrors.Errorf("taskPresenter.NotFound(): %w", err)
			}
			return ferrors.NewBusinessErrorf(err, "指定されたタスクが見つかりませんでした。")
		}
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	task.UnComplete()

	if err = u.taskRepository.Update(cctx, task, args.Version); err != nil {
		if errors.Is(err, repository.ErrNotAffected) {
			if err = u.taskPresenter.Conflict(cctx, "タスクは既に更新済みです。"); err != nil {
				return xerrors.Errorf("taskPresenter.Conflict(): %w", err)
			}
			return ferrors.NewBusinessErrorf(err, "タスクは既に更新済みです。")
		}
		return xerrors.Errorf("taskRepository.Update(): %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 task: %+v\n", task)

	taskUnCompleted := event.NewTaskUnCompleted(task)
	if err = u.taskEventRepository.Register(cctx, taskUnCompleted); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskUnCompleted: %+v\n", taskUnCompleted)

	u.publisher.Publish(taskUnCompleted)

	if err = u.taskPresenter.NilResponse(cctx); err != nil {
		return xerrors.Errorf("taskPresenter.NilResponse(): %w", err)
	}

	return nil
}
