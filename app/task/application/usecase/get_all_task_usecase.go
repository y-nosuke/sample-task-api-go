package usecase

import (
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"golang.org/x/xerrors"
)

type GetAllTaskUseCaseArgs struct {
}

type GetAllTaskUseCase struct {
	taskRepository repository.TaskRepository
	taskPresenter  presenter.TaskPresenter
}

func NewGetAllTaskUseCase(taskRepository repository.TaskRepository, taskPresenter presenter.TaskPresenter) *GetAllTaskUseCase {
	return &GetAllTaskUseCase{taskRepository, taskPresenter}
}

func (u *GetAllTaskUseCase) Invoke(cctx fcontext.Context, _ *GetAllTaskUseCaseArgs) error {
	tasks, err := u.taskRepository.GetAll(cctx)
	if err != nil {
		return xerrors.Errorf("taskRepository.GetAll(): %w", err)
	}

	if err = u.taskPresenter.TaskAllResponse(cctx, tasks); err != nil {
		return xerrors.Errorf("taskPresenter.TaskAllResponse(): %w", err)
	}

	return nil
}
