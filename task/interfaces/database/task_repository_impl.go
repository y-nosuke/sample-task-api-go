package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	fdatabase "github.com/y-nosuke/sample-task-api-go/framework/database/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/database/dao"
	"github.com/y-nosuke/sample-task-api-go/task/application/repositories"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
	"golang.org/x/xerrors"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() repositories.TaskRepository {
	return &TaskRepositoryImpl{}
}

func (t *TaskRepositoryImpl) Register(ctx context.Context, task *entities.Task) error {
	taskDto, err := taskDto(task)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if err = taskDto.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 taskDto: %+v\n", taskDto)

	task.CreatedAt = &taskDto.CreatedAt
	task.UpdatedAt = &taskDto.UpdatedAt
	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.Version = &version

	return nil
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]entities.Task, error) {
	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	taskDtos, err := dao.Tasks(qm.OrderBy("updated_at DESC")).All(ctx, tx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	fmt.Println("データベースからタスク一覧が取得されました。")

	tasks, err := tasks(taskDtos)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entities.Task, error) {
	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)

	taskId, err := id.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	taskDto, err := dao.FindTask(ctx, tx, taskId)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースからタスクが取得されました。 taskDto: %+v\n", taskDto)

	task, err := task(taskDto)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return task, nil
}

func (t *TaskRepositoryImpl) Update(ctx context.Context, task *entities.Task) error {
	taskDto, err := taskDto(task)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	oldVersion := taskDto.Version

	taskDto.Version, err = uuid.New().MarshalBinary()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if rowsAff, err := dao.Tasks(dao.TaskWhere.ID.EQ(taskDto.ID), dao.TaskWhere.Version.EQ(oldVersion)).UpdateAll(ctx, tx, dao.M{"id": taskDto.ID, "title": taskDto.Title, "detail": taskDto.Detail, "completed": taskDto.Completed, "deadline": taskDto.Deadline, "version": taskDto.Version}); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if rowsAff == 0 {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースのタスクが更新されました。 taskDto: %+v\n", taskDto)

	task.UpdatedAt = &taskDto.UpdatedAt
	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	task.Version = &version

	return nil
}

func (t *TaskRepositoryImpl) Delete(ctx context.Context, task *entities.Task) error {
	taskDto, err := taskDto(task)
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

func taskDto(task *entities.Task) (*dao.Task, error) {
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

	return &dao.Task{
		ID:        id,
		Title:     task.Title,
		Detail:    null.StringFromPtr(task.Detail),
		Completed: task.Completed,
		Deadline:  null.TimeFromPtr(task.Deadline),
		Version:   version,
	}, nil
}

func task(taskDto *dao.Task) (*entities.Task, error) {
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

	version, err := uuid.FromBytes(taskDto.Version)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &entities.Task{
		Id:        id,
		Title:     taskDto.Title,
		Detail:    detail,
		Completed: taskDto.Completed,
		Deadline:  deadline,
		CreatedAt: &taskDto.CreatedAt,
		UpdatedAt: &taskDto.UpdatedAt,
		Version:   &version,
	}, nil
}

func tasks(taskDtos []*dao.Task) ([]entities.Task, error) {
	var tasks []entities.Task
	for _, t := range taskDtos {
		task, err := task(t)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		tasks = append(tasks, *task)
	}

	return tasks, nil
}
