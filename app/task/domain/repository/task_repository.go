package repository

import (
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"golang.org/x/xerrors"
)

var (
	ErrNotFound    = xerrors.New("not found")
	ErrNotAffected = xerrors.New("not affected")
)

type TaskRepository interface {
	Register(fcontext.Context, *entity.Task) error
	GetAll(fcontext.Context) ([]*entity.Task, error)
	GetById(fcontext.Context, uuid.UUID) (*entity.Task, error)
	Update(fcontext.Context, *entity.Task, uuid.UUID) error
	Delete(fcontext.Context, *entity.Task) error
}
