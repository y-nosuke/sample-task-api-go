package repository

import (
	"context"
	"fmt"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure"
	fdatabase "github.com/y-nosuke/sample-task-api-go/app/framework/database/infrastructure"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	taskDto, err := taskDto(task, &a.UserId)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if err = taskDto.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 taskDto: %+v\n", taskDto)

	createdBy, err := uuid.FromBytes(taskDto.CreatedBy)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.CreatedBy = &createdBy
	updatedBy, err := uuid.FromBytes(taskDto.UpdatedBy)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.UpdatedBy = &updatedBy
	task.CreatedAt = &taskDto.CreatedAt
	task.UpdatedAt = &taskDto.UpdatedAt
	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.Version = &version

	return nil
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Task, error) {
	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	taskSlice, err := dao.Tasks(qm.OrderBy("updated_at DESC")).All(ctx, tx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	fmt.Println("データベースからタスク一覧が取得されました。")

	tasks, err := tasks(taskSlice)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)

	taskId, err := id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	taskDto, err := dao.FindTask(ctx, tx, taskId)
	if taskDto == nil {
		return nil, nil
	} else if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースからタスクが取得されました。 taskDto: %+v\n", taskDto)

	task, err := task(taskDto)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return task, nil
}

func (t *TaskRepositoryImpl) Update(ctx context.Context, task *entity.Task) (int, error) {
	a := ctx.Value(fauth.AUTH).(*auth.Auth)
	taskDto, err := taskDto(task, &a.UserId)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}

	oldVersion := taskDto.Version

	taskDto.Version, err = uuid.New().MarshalBinary()
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if rowsAff, err := dao.
		Tasks(dao.TaskWhere.ID.EQ(taskDto.ID), dao.TaskWhere.Version.EQ(oldVersion)).
		UpdateAll(ctx, tx, dao.M{"id": taskDto.ID, "title": taskDto.Title, "detail": taskDto.Detail, "completed": taskDto.Completed, "deadline": taskDto.Deadline, "version": taskDto.Version}); err != nil || rowsAff != 1 {
		return int(rowsAff), xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 taskDto: %+v\n", taskDto)

	updatedBy, err := uuid.FromBytes(taskDto.UpdatedBy)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}
	task.UpdatedBy = &updatedBy
	task.UpdatedAt = &taskDto.UpdatedAt
	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}
	task.Version = &version

	return 1, nil
}

func (t *TaskRepositoryImpl) Delete(ctx context.Context, task *entity.Task) error {
	a := ctx.Value(fauth.AUTH).(*auth.Auth)
	taskDto, err := taskDto(task, &a.UserId)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if _, err = taskDto.Delete(ctx, tx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースのタスクが削除されました。 taskDto: %+v\n", taskDto)

	return nil
}

func taskDto(task *entity.Task, userId *uuid.UUID) (*dao.Task, error) {
	id, err := task.Id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var version []byte
	if task.Version != nil {
		version, err = task.Version.MarshalBinary()
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
	}

	uid, err := userId.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &dao.Task{
		ID:        id,
		Title:     task.Title,
		Detail:    null.StringFromPtr(task.Detail),
		Completed: task.Completed,
		Deadline:  null.TimeFromPtr(task.Deadline),
		CreatedBy: uid,
		UpdatedBy: uid,
		Version:   version,
	}, nil
}

func task(taskDto *dao.Task) (*entity.Task, error) {
	id, err := uuid.FromBytes(taskDto.ID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var detail *string
	if taskDto.Detail.Valid {
		detail = &taskDto.Detail.String
	}

	var deadline *time.Time
	if taskDto.Deadline.Valid {
		deadline = &taskDto.Deadline.Time
	}

	createdBy, err := uuid.FromBytes(taskDto.CreatedBy)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	updatedBy, err := uuid.FromBytes(taskDto.UpdatedBy)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &entity.Task{
		Id:        &id,
		Title:     taskDto.Title,
		Detail:    detail,
		Completed: taskDto.Completed,
		Deadline:  deadline,
		CreatedBy: &createdBy,
		CreatedAt: &taskDto.CreatedAt,
		UpdatedBy: &updatedBy,
		UpdatedAt: &taskDto.UpdatedAt,
		Version:   &version,
	}, nil
}

func tasks(taskSlice dao.TaskSlice) ([]*entity.Task, error) {
	var tasks []*entity.Task
	for _, t := range taskSlice {
		task, err := task(t)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
