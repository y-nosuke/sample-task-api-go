package repositories

import (
	"context"

	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskRepository interface {
	Register(context.Context, *entities.Task) error
}
