package usecase

import (
	"fmt"
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type GetTaskUseCaseArgs struct {
	Id uuid.UUID
}

type GetTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewGetTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *GetTaskUseCase {
	return &GetTaskUseCase{taskRepository, taskPresenter}
}

func (u *GetTaskUseCase) Invoke(cctx fcontext.Context, args *GetTaskUseCaseArgs) error {
	fmt.Println("タスク取得処理を開始します。")

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

	fmt.Printf("データベースからタスクが取得されました。 task: %+v\n", task)

	if err = u.taskPresenter.GetTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.GetTaskResponse(): %w", err)
	}

	return nil
}
