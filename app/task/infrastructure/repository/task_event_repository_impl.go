package repository

import (
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/database"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"golang.org/x/xerrors"
)

type TaskEventRepositoryImpl struct {
}

func NewTaskEventRepositoryImpl() *TaskEventRepositoryImpl {
	return &TaskEventRepositoryImpl{}
}

func (t *TaskEventRepositoryImpl) Register(cctx fcontext.Context, taskEvent event.TaskEvent) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	a := auth.GetAuth(cctx)

	eTaskEvent, err := ETaskEvent(taskEvent, a.UserId)
	if err != nil {
		return xerrors.Errorf("mapping.RTask(): %w", err)
	}

	if err = eTaskEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskEvent.Insert(): %w", err)
	}

	createdBy, err := uuid.FromBytes(eTaskEvent.CreatedBy)
	if err != nil {
		return xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	taskEvent.Created(createdBy, eTaskEvent.CreatedAt)

	return nil
}
