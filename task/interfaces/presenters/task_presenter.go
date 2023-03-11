package presenters

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskPresenter struct {
}

func NewTaskPresenter() *TaskPresenter {
	return &TaskPresenter{}
}

func (p *TaskPresenter) TaskResponse(ectx echo.Context, task *entities.Task) error {
	return ectx.JSON(http.StatusOK, task)
}
