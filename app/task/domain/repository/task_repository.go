package repository

import (
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskRepository interface {
	Register(fcontext.Context, *entity.Task) error
	GetAll(fcontext.Context) ([]*entity.Task, error)
	GetById(fcontext.Context, uuid.UUID) (*entity.Task, error)
	Update(fcontext.Context, *entity.Task, uuid.UUID) (int, error)
	Delete(fcontext.Context, *entity.Task) error
}
