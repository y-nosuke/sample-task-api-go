package usecases

import (
	"context"

	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
	"golang.org/x/xerrors"
)

type RegisterTaskUseCaseInputData struct {
	Title    string
	Detail   string
	Deadline string
}

type RegisterTaskUseCaseOutputData struct {
	Task *entities.Task
}

type RegisterTaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewRegisterTaskUseCase(taskRepository repositories.TaskRepository) *RegisterTaskUseCase {
	return &RegisterTaskUseCase{taskRepository}
}

func (u *RegisterTaskUseCase) Invoke(ctx context.Context, inputData *RegisterTaskUseCaseInputData) (*RegisterTaskUseCaseOutputData, error) {
	task, err := entities.NewTask(inputData.Title, inputData.Detail, inputData.Deadline)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	err = u.taskRepository.Register(ctx, task)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &RegisterTaskUseCaseOutputData{task}, nil
}
