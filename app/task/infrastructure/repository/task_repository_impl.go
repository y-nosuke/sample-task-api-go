package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
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

func (t *TaskRepositoryImpl) Register(cctx fcontext.Context, task *entity.Task) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	a := auth.GetAuth(cctx)

	newVersion := uuid.New()
	rTask, err := RTask(task, a.UserId, newVersion)
	if err != nil {
		return xerrors.Errorf("mapping.RTask(): %w", err)
	}

	if err = rTask.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf("rTask.Insert(): %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 rTask: %+v\n", rTask)

	createdBy, err := uuid.FromBytes(rTask.CreatedBy)
	if err != nil {
		return xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.CreatedBy = createdBy
	task.CreatedAt = rTask.CreatedAt

	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.UpdatedBy = updatedBy
	task.UpdatedAt = rTask.UpdatedAt

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

	fmt.Println("データベースからタスク一覧が取得されました。")

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

func (t *TaskRepositoryImpl) Update(cctx fcontext.Context, task *entity.Task, version uuid.UUID) (int, error) {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)
	a := auth.GetAuth(cctx)

	newVersion := uuid.New()
	rTask, err := RTask(task, a.UserId, newVersion)
	if err != nil {
		return 0, xerrors.Errorf("mapping.RTask(): %w", err)
	}

	byteVersion, err := version.MarshalBinary()
	if err != nil {
		return 0, xerrors.Errorf("version.MarshalBinary(): %w", err)
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
	rowsAff, err := dao.RTasks(qs...).UpdateAll(ctx, tx, cols)
	if err != nil {
		return 0, xerrors.Errorf("dao.RTasks().UpdateAll(): %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 rTask: %+v\n", rTask)

	updatedBy, err := uuid.FromBytes(rTask.UpdatedBy)
	if err != nil {
		return 0, xerrors.Errorf("uuid.FromBytes(): %w", err)
	}
	task.UpdatedBy = updatedBy
	task.UpdatedAt = rTask.UpdatedAt

	task.Version = newVersion

	return int(rowsAff), nil
}

func (t *TaskRepositoryImpl) Delete(cctx fcontext.Context, task *entity.Task) error {
	ctx := cctx.GetContext()
	tx := database.GetTransaction(cctx)

	a := auth.GetAuth(cctx)
	rTask, err := RTask(task, a.UserId, task.Version)
	if err != nil {
		return xerrors.Errorf("RTask(): %w", err)
	}

	if _, err = rTask.Delete(ctx, tx); err != nil {
		return xerrors.Errorf("rTask.Delete(): %w", err)
	}

	fmt.Printf("データベースのタスクが削除されました。 rTask: %+v\n", rTask)

	return nil
}
