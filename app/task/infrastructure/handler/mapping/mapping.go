package mapping

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"time"
)

func RegisterTaskUseCaseArgs(request *openapi.RegisterTaskRequest) *usecase.RegisterTaskUseCaseArgs {
	var deadline *time.Time
	if request.Deadline != nil {
		deadline = &request.Deadline.Time
	}

	return &usecase.RegisterTaskUseCaseArgs{
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: deadline,
	}
}

func UpdateTaskUseCaseArgs(id uuid.UUID, request *openapi.UpdateTaskRequest) (*usecase.UpdateTaskUseCaseArgs, error) {
	return &usecase.UpdateTaskUseCaseArgs{
		Id:       id,
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: &request.Deadline.Time,
		Version:  &request.Version,
	}, nil
}

func CompleteTaskUseCaseArgs(id uuid.UUID, request *openapi.CompleteTaskRequest) (*usecase.CompleteTaskUseCaseArgs, error) {
	return &usecase.CompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}

func UnCompleteTaskUseCaseArgs(id uuid.UUID, request *openapi.UnCompleteTaskRequest) (*usecase.UnCompleteTaskUseCaseArgs, error) {
	return &usecase.UnCompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}
