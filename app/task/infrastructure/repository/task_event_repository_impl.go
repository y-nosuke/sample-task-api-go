package repository

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/database"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"golang.org/x/xerrors"
)

type TaskEventRepositoryImpl struct{}

func NewTaskEventRepositoryImpl() *TaskEventRepositoryImpl {
	return &TaskEventRepositoryImpl{}
}

func (t *TaskEventRepositoryImpl) RegisterTaskCreated(cctx fcontext.Context, taskCreated *event.TaskCreated) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	eTaskEvent, eTaskCreated, err := ETaskCreated(taskCreated, userId)
	if err != nil {
		return xerrors.Errorf("ETaskCreated(): %w", err)
	}

	if err = eTaskEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskEvent.Insert(): %w", err)
	}

	if err = eTaskCreated.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskCreated.Insert(): %w", err)
	}

	return nil
}

func (t *TaskEventRepositoryImpl) RegisterTaskUpdated(cctx fcontext.Context, taskUpdated *event.TaskEdited) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	eTaskEvent, eTaskCreated, err := ETaskUpdated(taskUpdated, userId)
	if err != nil {
		return xerrors.Errorf("ETaskCreated(): %w", err)
	}

	if err = eTaskEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskEvent.Insert(): %w", err)
	}

	if err = eTaskCreated.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskCreated.Insert(): %w", err)
	}

	return nil
}

func (t *TaskEventRepositoryImpl) RegisterTaskCompleted(cctx fcontext.Context, taskCompleted *event.TaskCompleted) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	eTaskCompleted, err := ETaskCompleted(taskCompleted, userId)
	if err != nil {
		return xerrors.Errorf("ETaskCompleted(): %w", err)
	}

	if err = eTaskCompleted.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskCompleted.Insert(): %w", err)
	}

	return nil
}

func (t *TaskEventRepositoryImpl) RegisterTaskUnCompleted(cctx fcontext.Context, taskUnCompleted *event.TaskUnCompleted) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	eTaskUnCompleted, err := ETaskUnCompleted(taskUnCompleted, userId)
	if err != nil {
		return xerrors.Errorf("ETaskUnCompleted(): %w", err)
	}

	if err = eTaskUnCompleted.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskUnCompleted.Insert(): %w", err)
	}

	return nil
}

func (t *TaskEventRepositoryImpl) RegisterTaskDeleted(cctx fcontext.Context, taskDeleted *event.TaskDeleted) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	eTaskDeleted, err := ETaskDeleted(taskDeleted, userId)
	if err != nil {
		return xerrors.Errorf("ETaskDeleted(): %w", err)
	}

	if err = eTaskDeleted.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("eTaskDeleted.Insert(): %w", err)
	}

	return nil
}
