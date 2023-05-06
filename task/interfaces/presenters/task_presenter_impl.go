package presenters

import (
	"context"
	"net/http"

	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
	"github.com/y-nosuke/sample-task-api-go/task/enterprise/entities"
)

type TaskPresenter struct {
}

func NewTaskPresenter() *TaskPresenter {
	return &TaskPresenter{}
}

func (p *TaskPresenter) RegisterTaskResponse(ctx context.Context, task *entities.Task) error {
	return p.taskResponse(ctx, http.StatusCreated, task)
}

func (p *TaskPresenter) UpdateTaskResponse(ctx context.Context, task *entities.Task) error {
	return p.taskResponse(ctx, http.StatusOK, task)
}

func (p *TaskPresenter) GetTaskResponse(ctx context.Context, task *entities.Task) error {
	return p.taskResponse(ctx, http.StatusOK, task)
}

func (p *TaskPresenter) taskResponse(ctx context.Context, code int, task *entities.Task) error {
	ectx := fcontext.Ectx(ctx)
	taskForm := openapi.TaskResponse{
		TaskForm: *taskForm(task),
	}

	return ectx.JSON(code, &taskForm)
}

func (p *TaskPresenter) TaskAllResponse(ctx context.Context, tasks []*entities.Task) error {
	ectx := fcontext.Ectx(ctx)

	taskForm := openapi.GetAllTasksResponse{
		TaskForms: taskForms(tasks),
	}

	return ectx.JSON(http.StatusOK, &taskForm)
}

func (p *TaskPresenter) NilResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)

	return ectx.NoContent(http.StatusOK)
}

func (p *TaskPresenter) NoContentResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)

	return ectx.NoContent(http.StatusNoContent)
}

func taskForm(task *entities.Task) *openapi.TaskForm {
	var deadline *openapiTypes.Date
	if task.Deadline != nil {
		deadline = &openapiTypes.Date{Time: *task.Deadline}
	}

	return &openapi.TaskForm{
		Id:        task.Id,
		Title:     task.Title,
		Detail:    task.Detail,
		Completed: &task.Completed,
		Deadline:  deadline,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		Version:   *task.Version,
	}
}

func taskForms(tasks []*entities.Task) []openapi.TaskForm {
	taskForms := make([]openapi.TaskForm, 0, 10)
	for _, t := range tasks {
		taskForms = append(taskForms, *taskForm(t))
	}

	return taskForms
}
