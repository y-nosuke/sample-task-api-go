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
	fmt.Println("タスク更新処理を開始します。")

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

	task.Update(args.Title, args.Detail, args.Deadline)

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

	taskUpdated := event.NewTaskUpdated(task)
	if err = u.taskEventRepository.Register(cctx, taskUpdated); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskUpdated: %+v\n", taskUpdated)

	u.publisher.Publish(taskUpdated)

	if err = u.taskPresenter.UpdateTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.UpdateTaskResponse(): %w", err)
	}

	return nil
}
