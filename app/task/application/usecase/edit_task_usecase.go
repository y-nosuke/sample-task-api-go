package usecase

import (
	"errors"
	"fmt"
	"time"

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

type EditTaskUseCaseArgs struct {
	Id        uuid.UUID
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
	Version   uuid.UUID
}

type EditTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewEditTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *EditTaskUseCase {
	return &EditTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *EditTaskUseCase) Invoke(cctx fcontext.Context, args *EditTaskUseCaseArgs) error {
	fmt.Println("タスク編集処理を開始します。")

	task, err := u.taskRepository.GetById(cctx, args.Id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ferrors.NewNotFoundErrorf(err, "指定されたタスクが見つかりませんでした。")
		}
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	taskEdited, err := task.Edit(args.Title, args.Detail, args.Completed, args.Deadline, auth.GetUserId(cctx))
	if err != nil {
		return xerrors.Errorf("task.Edit(): %w", err)
	}

	if err = u.taskRepository.Update(cctx, task, args.Version); err != nil {
		if errors.Is(err, repository.ErrNotAffected) {
			return ferrors.NewConflictErrorf(err, "タスクは既に更新済みです。")
		}
		return xerrors.Errorf("taskRepository.Edit(): %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 task: %+v\n", task)

	if err = u.taskEventRepository.RegisterTaskEdited(cctx, taskEdited); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskEdited: %+v\n", taskEdited)

	u.publisher.Publish(taskEdited)

	if err = u.taskPresenter.UpdateTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.UpdateTaskResponse(): %w", err)
	}

	return nil
}
