package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure"
	fdatabase "github.com/y-nosuke/sample-task-api-go/app/framework/database/infrastructure"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/repository/mapping"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepositoryImpl() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{}
}

func (t *TaskRepositoryImpl) Register(ctx context.Context, task *entity.Task) error {
	a := ctx.Value(fauth.AUTH).(*auth.Auth)
	rTask, err := mapping.RTask(task, &a.UserId)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if err = rTask.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 rTask: %+v\n", rTask)

	createdBy, err := uuid.FromBytes(rTask.CreatedBy)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.CreatedBy = &createdBy
	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.UpdatedBy = &updatedBy
	task.CreatedAt = &rTask.CreatedAt
	task.UpdatedAt = &rTask.UpdatedAt
	version, err := uuid.FromBytes(rTask.Version)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.Version = &version

	return nil
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Task, error) {
	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	rTaskSlice, err := dao.RTasks(qm.OrderBy("updated_at DESC")).All(ctx, tx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	fmt.Println("データベースからタスク一覧が取得されました。")

	taskSlice, err := mapping.TaskSlice(rTaskSlice)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return taskSlice, nil
}

func (t *TaskRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)

	taskId, err := id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	rTask, err := dao.FindRTask(ctx, tx, taskId)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if rTask == nil {
		return nil, nil
	}

	fmt.Printf("データベースからタスクが取得されました。 rTask: %+v\n", rTask)

	task, err := mapping.Task(rTask)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return task, nil
}

func (t *TaskRepositoryImpl) Update(ctx context.Context, task *entity.Task) (int, error) {
	a := ctx.Value(fauth.AUTH).(*auth.Auth)
	rTask, err := mapping.RTask(task, &a.UserId)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}

	oldVersion := rTask.Version

	rTask.Version, err = uuid.New().MarshalBinary()
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if rowsAff, err := dao.
		RTasks(dao.RTaskWhere.ID.EQ(rTask.ID), dao.RTaskWhere.Version.EQ(oldVersion)).
		UpdateAll(ctx, tx, dao.M{"id": rTask.ID, "title": rTask.Title, "detail": rTask.Detail, "completed": rTask.Completed, "deadline": rTask.Deadline, "version": rTask.Version}); err != nil || rowsAff != 1 {
		return int(rowsAff), xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 rTask: %+v\n", rTask)

	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}
	task.UpdatedBy = &updatedBy
	task.UpdatedAt = &rTask.UpdatedAt
	version, err := uuid.FromBytes(rTask.Version)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}
	task.Version = &version

	return 1, nil
}

func (t *TaskRepositoryImpl) Delete(ctx context.Context, task *entity.Task) error {
	a := ctx.Value(fauth.AUTH).(*auth.Auth)
	rTask, err := mapping.RTask(task, &a.UserId)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if _, err = rTask.Delete(ctx, tx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースのタスクが削除されました。 rTask: %+v\n", rTask)

	return nil
}
