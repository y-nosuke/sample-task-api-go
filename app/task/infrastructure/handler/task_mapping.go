package handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

func RegisterTaskUseCaseArgs(request *openapi.CreateTaskRequest) *usecase.CreateTaskUseCaseArgs {
	var deadline *time.Time
	if request.Deadline != nil {
		deadline = &request.Deadline.Time
	}

	return &usecase.CreateTaskUseCaseArgs{
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: deadline,
	}
}

func EditTaskUseCaseArgs(id uuid.UUID, request *openapi.EditTaskRequest) *usecase.EditTaskUseCaseArgs {
	var deadline *time.Time
	if request.Deadline != nil {
		deadline = &request.Deadline.Time
	}
	var completed bool
	if request.Completed != nil {
		completed = *request.Completed
	}

	return &usecase.EditTaskUseCaseArgs{
		Id:        id,
		Title:     request.Title,
		Detail:    request.Detail,
		Completed: completed,
		Deadline:  deadline,
		Version:   request.Version,
	}
}
