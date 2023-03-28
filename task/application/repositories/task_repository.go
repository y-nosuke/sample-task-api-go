package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskRepository interface {
	Register(context.Context, *entities.Task) error
	GetAll(context.Context) ([]entities.Task, error)
	GetById(context.Context, uuid.UUID) (*entities.Task, error)
	Update(context.Context, *entities.Task) (int, error)
	Delete(context.Context, *entities.Task) error
}
