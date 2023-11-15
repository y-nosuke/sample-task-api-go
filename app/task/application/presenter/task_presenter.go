package presenter

import (
	"context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskPresenter interface {
	RegisterTaskResponse(context.Context, *entity.Task) error
	UpdateTaskResponse(context.Context, *entity.Task) error
	GetTaskResponse(context.Context, *entity.Task) error
	TaskAllResponse(context.Context, entity.TaskSlice) error
	NilResponse(context.Context) error
	NoContentResponse(context.Context) error

	BadRequest(context.Context, string, error) error
	Forbidden(context.Context, string) error
	NotFound(context.Context, string) error
	Conflict(context.Context, string) error
	InternalServerError(context.Context, string) error
}
