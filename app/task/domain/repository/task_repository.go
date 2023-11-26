package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskRepository interface {
	Register(context.Context, *entity.Task) error
	GetAll(context.Context) ([]*entity.Task, error)
	GetById(context.Context, uuid.UUID) (*entity.Task, error)
	Update(context.Context, *entity.Task, *uuid.UUID) (int, error)
	Delete(context.Context, *entity.Task) error
}
