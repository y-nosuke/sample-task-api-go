package presenter

import (
	"context"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"net/http"

	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

type TaskPresenterImpl struct {
}

func NewTaskPresenterImpl() *TaskPresenterImpl {
	return &TaskPresenterImpl{}
}

func (p *TaskPresenterImpl) RegisterTaskResponse(ctx context.Context, task *entity.Task) error {
	return p.taskResponse(ctx, http.StatusCreated, task)
}

func (p *TaskPresenterImpl) UpdateTaskResponse(ctx context.Context, task *entity.Task) error {
	return p.taskResponse(ctx, http.StatusOK, task)
}

func (p *TaskPresenterImpl) GetTaskResponse(ctx context.Context, task *entity.Task) error {
	return p.taskResponse(ctx, http.StatusOK, task)
}

func (p *TaskPresenterImpl) taskResponse(ctx context.Context, code int, task *entity.Task) error {
	ectx := fcontext.Ectx(ctx)
	taskForm := openapi.TaskResponse{
		TaskForm: *taskForm(task),
	}

	return ectx.JSON(code, &taskForm)
}

func (p *TaskPresenterImpl) TaskAllResponse(ctx context.Context, tasks []*entity.Task) error {
	ectx := fcontext.Ectx(ctx)

	taskForm := openapi.GetAllTasksResponse{
		TaskForms: taskForms(tasks),
	}

	return ectx.JSON(http.StatusOK, &taskForm)
}

func (p *TaskPresenterImpl) NilResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)

	return ectx.NoContent(http.StatusOK)
}

func (p *TaskPresenterImpl) NoContentResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)

	return ectx.NoContent(http.StatusNoContent)
}

func taskForm(task *entity.Task) *openapi.TaskForm {
	var deadline *openapi.NullableDeadline
	if task.Deadline != nil {
		deadline = &openapi.NullableDeadline{Time: *task.Deadline}
	}

	return &openapi.TaskForm{
		Id:        *task.Id,
		Title:     task.Title,
		Detail:    task.Detail,
		Completed: &task.Completed,
		Deadline:  deadline,
		CreatedBy: *task.CreatedBy,
		CreatedAt: *task.CreatedAt,
		UpdatedBy: *task.UpdatedBy,
		UpdatedAt: *task.UpdatedAt,
		Version:   *task.Version,
	}
}

func taskForms(tasks []*entity.Task) []openapi.TaskForm {
	taskForms := make([]openapi.TaskForm, 0, 10)
	for _, t := range tasks {
		taskForms = append(taskForms, *taskForm(t))
	}

	return taskForms
}
