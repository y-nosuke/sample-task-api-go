package handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
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

func UpdateTaskUseCaseArgs(id uuid.UUID, request *openapi.UpdateTaskRequest) *usecase.UpdateTaskUseCaseArgs {
	return &usecase.UpdateTaskUseCaseArgs{
		Id:       id,
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: &request.Deadline.Time,
		Version:  &request.Version,
	}
}

func CompleteTaskUseCaseArgs(id uuid.UUID, request *openapi.CompleteTaskRequest) *usecase.CompleteTaskUseCaseArgs {
	return &usecase.CompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}
}

func UnCompleteTaskUseCaseArgs(id uuid.UUID, request *openapi.UnCompleteTaskRequest) *usecase.UnCompleteTaskUseCaseArgs {
	return &usecase.UnCompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}
}
