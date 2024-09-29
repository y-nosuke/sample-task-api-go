package presenter

import (
	"context"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	fpresenter "github.com/y-nosuke/sample-task-api-go/app/framework/io/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"net/http"
)

type TaskPresenterImpl struct {
	fpresenter.BusinessErrorPresenter
}

func NewTaskPresenterImpl(businessErrorPresenter fpresenter.BusinessErrorPresenter) *TaskPresenterImpl {
	return &TaskPresenterImpl{businessErrorPresenter}
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
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(code, TaskResponse(task))
}

func (p *TaskPresenterImpl) TaskAllResponse(ctx context.Context, taskSlice entity.TaskSlice) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusOK, GetAllTasksResponse(taskSlice))
}

func (p *TaskPresenterImpl) NilResponse(ctx context.Context) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.NoContent(http.StatusOK)
}

func (p *TaskPresenterImpl) NoContentResponse(ctx context.Context) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.NoContent(http.StatusNoContent)
}
