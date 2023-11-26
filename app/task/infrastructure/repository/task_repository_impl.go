package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	"github.com/y-nosuke/sample-task-api-go/app/framework/database"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepositoryImpl() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{}
}

func (t *TaskRepositoryImpl) Register(ctx context.Context, task *entity.Task) error {
	a := auth.GetAuth(ctx)
	rTask, err := RTask(task, &a.UserId, task.Version)
	if err != nil {
		return xerrors.Errorf("mapping.RTask(): %w", err)
	}

	tx := database.GetTransaction(ctx)
	if err = rTask.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("rTask.Insert(): %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 rTask: %+v\n", rTask)

	createdBy, err := uuid.FromBytes(rTask.CreatedBy)
	if err != nil {
		return xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.CreatedBy = &createdBy
	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.UpdatedBy = &updatedBy
	task.CreatedAt = &rTask.CreatedAt
	task.UpdatedAt = &rTask.UpdatedAt
	version, err := uuid.FromBytes(rTask.Version)
	if err != nil {
		return xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.Version = &version

	return nil
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Task, error) {
	tx := database.GetTransaction(ctx)
	rTaskSlice, err := dao.RTasks(qm.OrderBy("updated_at DESC")).All(ctx, tx)
	if err != nil {
		return nil, xerrors.Errorf("dao.RTasks(): %w", err)
	}

	fmt.Println("データベースからタスク一覧が取得されました。")

	taskSlice, err := TaskSlice(rTaskSlice)
	if err != nil {
		return nil, xerrors.Errorf("mapping.TaskSlice(): %w", err)
	}

	return taskSlice, nil
}

func (t *TaskRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	taskId, err := id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("id.MarshalBinary(): %w", err)
	}

	tx := database.GetTransaction(ctx)
	rTask, err := dao.FindRTask(ctx, tx, taskId)
	if err != nil {
		return nil, xerrors.Errorf("dao.FindRTask(): %w", err)
	}

	if rTask == nil {
		return nil, nil
	}

	fmt.Printf("データベースからタスクが取得されました。 rTask: %+v\n", rTask)

	task, err := Task(rTask)
	if err != nil {
		return nil, xerrors.Errorf("mapping.Task(): %w", err)
	}

	return task, nil
}

func (t *TaskRepositoryImpl) Update(ctx context.Context, task *entity.Task, oldVersion *uuid.UUID) (int, error) {
	a := auth.GetAuth(ctx)
	newVersion := uuid.New()
	rTask, err := RTask(task, &a.UserId, &newVersion)
	if err != nil {
		return 0, xerrors.Errorf("mapping.RTask(): %w", err)
	}

	byteOldVersion, err := oldVersion.MarshalBinary()
	if err != nil {
		return 0, xerrors.Errorf("oldVersion.MarshalBinary(): %w", err)
	}

	tx := database.GetTransaction(ctx)
	rowsAff, err := dao.
		RTasks(dao.RTaskWhere.ID.EQ(rTask.ID), dao.RTaskWhere.Version.EQ(byteOldVersion)).
		UpdateAll(ctx, tx, dao.M{"id": rTask.ID, "title": rTask.Title, "detail": rTask.Detail, "completed": rTask.Completed, "deadline": rTask.Deadline, "version": rTask.Version})
	if err != nil {
		return 0, xerrors.Errorf("dao.RTasks().UpdateAll(): %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 rTask: %+v\n", rTask)

	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return 0, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.UpdatedBy = &updatedBy
	task.UpdatedAt = &rTask.UpdatedAt
	_version, err := uuid.FromBytes(rTask.Version)
	if err != nil {
		return 0, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.Version = &_version

	return int(rowsAff), nil
}

func (t *TaskRepositoryImpl) Delete(ctx context.Context, task *entity.Task) error {
	a := auth.GetAuth(ctx)
	rTask, err := RTask(task, &a.UserId, task.Version)
	if err != nil {
		return xerrors.Errorf("mapping.RTask(): %w", err)
	}

	tx := database.GetTransaction(ctx)
	if _, err = rTask.Delete(ctx, tx); err != nil {
		return xerrors.Errorf("rTask.Delete(): %w", err)
	}

	fmt.Printf("データベースのタスクが削除されました。 rTask: %+v\n", rTask)

	return nil
}
