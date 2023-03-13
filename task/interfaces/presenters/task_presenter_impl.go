package presenters

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskPresenter struct {
}

func NewTaskPresenter() *TaskPresenter {
	return &TaskPresenter{}
}

func (p *TaskPresenter) TaskResponse(ctx context.Context, task *entities.Task) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusOK, task)
}
