package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	id, err := task.Id.MarshalBinary()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	taskDto := dao.Task{
		ID:        id,
		Title:     task.Title,
		Detail:    null.StringFrom(task.Detail),
		Completed: task.Completed,
		Deadline:  null.TimeFrom(task.Deadline),
	}

	tx := ctx.Value(fdatabase.TRANSACTION).(boil.ContextExecutor)
	if err = taskDto.Insert(ctx, tx, boil.Infer()); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Printf("データベースにタスクが登録されました。 taskDto: %+v\n", taskDto)

	task.CreatedAt = taskDto.CreatedAt
	task.UpdatedAt = taskDto.UpdatedAt
	task.Version, err = uuid.FromBytes(taskDto.Version)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
