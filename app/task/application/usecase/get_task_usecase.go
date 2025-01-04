package usecase

import (
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
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
		if errors.Is(err, repository.ErrNotFound) {
			return ferrors.NewNotFoundErrorf(err, "指定されたタスクが見つかりませんでした。")
		}
		return xerrors.Errorf("taskRepository.GetById(): %w", err)
	}

	fmt.Printf("データベースからタスクが取得されました。 task: %+v\n", task)

	if err = u.taskPresenter.GetTaskResponse(cctx, task); err != nil {
		return xerrors.Errorf("taskPresenter.GetTaskResponse(): %w", err)
	}

	return nil
}
