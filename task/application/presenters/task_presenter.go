package presenters

import (
	"context"

	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskPresenter interface {
	RegisterTaskResponse(context.Context, *entities.Task) error
	UpdateTaskResponse(context.Context, *entities.Task) error
	GetTaskResponse(context.Context, *entities.Task) error
	TaskAllResponse(context.Context, []*entities.Task) error
	NilResponse(context.Context) error
}
