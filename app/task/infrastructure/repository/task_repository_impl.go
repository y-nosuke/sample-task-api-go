package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/database"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/repository"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepositoryImpl() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{}
}

func (t *TaskRepositoryImpl) Register(cctx fcontext.Context, task *entity.Task) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	newVersion := uuid.New()
	rTask, err := RTask(task, userId, newVersion, true)
	if err != nil {
		return xerrors.Errorf("RTask(): %w", err)
	}

	if err = rTask.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("rTask.Insert(): %w", err)
	}

	task.Version = newVersion

	return nil
}

func (t *TaskRepositoryImpl) GetAll(cctx fcontext.Context) ([]*entity.Task, error) {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)

	qs := []qm.QueryMod{
		qm.OrderBy(dao.RTaskColumns.UpdatedAt + " DESC"),
	}
	rTaskSlice, err := dao.RTasks(qs...).All(ctx, tx)
	if err != nil {
		return nil, xerrors.Errorf("dao.RTasks(): %w", err)
	}

	taskSlice, err := TaskSlice(rTaskSlice)
	if err != nil {
		return nil, xerrors.Errorf("TaskSlice(): %w", err)
	}

	return taskSlice, nil
}

func (t *TaskRepositoryImpl) GetById(cctx fcontext.Context, id uuid.UUID) (*entity.Task, error) {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)

	taskId, err := id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("id.MarshalBinary(): %w", err)
	}

	rTask, err := dao.FindRTask(ctx, tx, taskId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, xerrors.Errorf("dao.FindRTask(): %w", err)
	}

	task, err := Task(rTask)
	if err != nil {
		return nil, xerrors.Errorf("Task(): %w", err)
	}

	return task, nil
}

func (t *TaskRepositoryImpl) Update(cctx fcontext.Context, task *entity.Task, version uuid.UUID) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	newVersion := uuid.New()
	rTask, err := RTask(task, userId, newVersion, false)
	if err != nil {
		return xerrors.Errorf("RTask(): %w", err)
	}

	byteVersion, err := version.MarshalBinary()
	if err != nil {
		return xerrors.Errorf("version.MarshalBinary(): %w", err)
	}

	qs := []qm.QueryMod{
		dao.RTaskWhere.ID.EQ(rTask.ID),
		dao.RTaskWhere.Version.EQ(byteVersion),
	}
	cols := dao.M{
		dao.RTaskColumns.ID:        rTask.ID,
		dao.RTaskColumns.Title:     rTask.Title,
		dao.RTaskColumns.Detail:    rTask.Detail,
		dao.RTaskColumns.Completed: rTask.Completed,
		dao.RTaskColumns.Deadline:  rTask.Deadline,
		dao.RTaskColumns.Version:   rTask.Version,
	}
	if rowsAff, err := dao.RTasks(qs...).UpdateAll(ctx, tx, cols); err != nil {
		return xerrors.Errorf("dao.RTasks().UpdateAll(): %w", err)
	} else if rowsAff == 0 {
		return repository.ErrNotAffected
	}

	task.Version = newVersion

	return nil
}

func (t *TaskRepositoryImpl) Delete(cctx fcontext.Context, task *entity.Task) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	userId := auth.GetUserId(cctx)

	rTask, err := RTask(task, userId, task.Version, false)
	if err != nil {
		return xerrors.Errorf("RTask(): %w", err)
	}

	if rowsAff, err := rTask.Delete(ctx, tx); err != nil {
		return xerrors.Errorf("rTask.Delete(): %w", err)
	} else if rowsAff == 0 {
		return repository.ErrNotAffected
	}

	return nil
}
