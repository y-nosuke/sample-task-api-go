package presenter

import (
	"context"

	fpresenter "github.com/y-nosuke/sample-task-api-go/app/framework/io/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskPresenter interface {
	RegisterTaskResponse(context.Context, *entity.Task) error
	UpdateTaskResponse(context.Context, *entity.Task) error
	GetTaskResponse(context.Context, *entity.Task) error
	TaskAllResponse(context.Context, entity.TaskSlice) error
	NilResponse(context.Context) error
	NoContentResponse(context.Context) error
	fpresenter.BusinessErrorPresenter
}
