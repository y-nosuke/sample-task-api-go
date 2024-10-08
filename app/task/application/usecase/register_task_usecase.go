package usecase

import (
	"fmt"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	nevent "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/factory"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
	"time"
)

type RegisterTaskUseCaseArgs struct {
	Title    string
	Detail   *string
	Deadline *time.Time
}

type RegisterTaskUseCase struct {
	taskRepository      repository.TaskRepository
	taskEventRepository repository.TaskEventRepository
	taskPresenter       presenter.TaskPresenter
	publisher           observer.Publisher[nevent.DomainEvent]
}

func NewRegisterTaskUseCase(taskRepository repository.TaskRepository, taskEventRepository repository.TaskEventRepository, taskPresenter presenter.TaskPresenter, publisher observer.Publisher[nevent.DomainEvent]) *RegisterTaskUseCase {
	return &RegisterTaskUseCase{taskRepository, taskEventRepository, taskPresenter, publisher}
}

func (u *RegisterTaskUseCase) Invoke(cctx fcontext.Context, args *RegisterTaskUseCaseArgs) error {
	fmt.Println("タスク登録処理を開始します。")

	task := factory.CreateTask(args.Title, args.Detail, args.Deadline)
	if err := u.taskRepository.Register(cctx, task); err != nil {
		return xerrors.Errorf("taskRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 task: %+v\n", task)

	taskCreated := event.NewTaskCreated(task)
	if err := u.taskEventRepository.Register(cctx, taskCreated); err != nil {
		return xerrors.Errorf("taskEventRepository.Register(): %w", err)
	}

	fmt.Printf("データベースにタスクイベントが登録されました。 taskCreated: %+v\n", taskCreated)

	u.publisher.Publish(taskCreated)

	if err := u.taskPresenter.RegisterTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.RegisterTaskResponse(): %w", err)
	}

	return nil
}
