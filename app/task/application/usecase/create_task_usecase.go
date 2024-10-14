package usecase

import (
	"fmt"
	"time"

	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/factory"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type CreateTaskUseCaseArgs struct {
	Title    string
	Detail   *string
	Deadline *time.Time
}

type CreateTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewCreateTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *CreateTaskUseCase {
	return &CreateTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *CreateTaskUseCase) Invoke(cctx fcontext.Context, args *CreateTaskUseCaseArgs) error {
	fmt.Println("タスク登録処理を開始します。")

	task, taskCreated, err := factory.CreateTask(args.Title, args.Detail, args.Deadline, auth.GetUserId(cctx))
	if err != nil {
		return xerrors.Errorf("factory.CreateTask(): %w", err)
	}

	if err = u.taskRepository.Register(cctx, task); err != nil {
		return xerrors.Errorf("taskRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 task: %+v\n", task)

	if err = u.taskEventRepository.RegisterTaskCreated(cctx, taskCreated); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskCreated: %+v\n", taskCreated)

	u.publisher.Publish(taskCreated)

	if err = u.taskPresenter.RegisterTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.RegisterTaskResponse(): %w", err)
	}

	return nil
}
