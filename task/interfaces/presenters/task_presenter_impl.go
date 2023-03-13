package presenters

import (
	"context"
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskPresenter struct {
}

func NewTaskPresenter() *TaskPresenter {
	return &TaskPresenter{}
}

func (p *TaskPresenter) TaskResponse(ctx context.Context, task *entities.Task) error {
	ectx := fcontext.Ectx(ctx)
	taskForm := openapi.TaskResponse{
		TaskForm: openapi.TaskForm{
			Id:        task.Id,
			Title:     task.Title,
			Detail:    &task.Detail,
			Status:    &task.Completed,
			Deadline:  &openapi_types.Date{Time: task.Deadline},
			CreatedAt: &task.CreatedAt,
			UpdatedAt: &task.UpdatedAt,
			Version:   task.Version,
		},
	}
	return ectx.JSON(http.StatusOK, &taskForm)
}
