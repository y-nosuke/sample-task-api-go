package presenters

import (
	"context"

	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskPresenter interface {
	TaskResponse(context.Context, *entities.Task) error
}
