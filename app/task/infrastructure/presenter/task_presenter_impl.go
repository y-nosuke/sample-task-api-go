package presenter

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
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

func (p *TaskPresenterImpl) BadRequest(ctx context.Context, message string, err error) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusBadRequest, &openapi.ErrorResponse{Message: BadRequestMessage(message, err)})
}

func (p *TaskPresenterImpl) Forbidden(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusForbidden, &openapi.ErrorResponse{Message: message})
}

func (p *TaskPresenterImpl) NotFound(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusNotFound, &openapi.ErrorResponse{Message: message})
}

func (p *TaskPresenterImpl) Conflict(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusConflict, &openapi.ErrorResponse{Message: message})
}

func (p *TaskPresenterImpl) InternalServerError(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusInternalServerError, &openapi.ErrorResponse{Message: message})
}
