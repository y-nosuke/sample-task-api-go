package presenter

import (
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskPresenter interface {
	RegisterTaskResponse(fcontext.Context, *entity.Task) error
	UpdateTaskResponse(fcontext.Context, *entity.Task) error
	GetTaskResponse(fcontext.Context, *entity.Task) error
	TaskAllResponse(fcontext.Context, entity.TaskSlice) error
	NilResponse(fcontext.Context) error
	NoContentResponse(fcontext.Context) error
}
