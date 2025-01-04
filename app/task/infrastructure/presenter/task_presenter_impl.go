package presenter

import (
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskPresenterImpl struct{}

func NewTaskPresenterImpl() *TaskPresenterImpl {
	return &TaskPresenterImpl{}
}

func (p *TaskPresenterImpl) RegisterTaskResponse(cctx fcontext.Context, task *entity.Task) error {
	return p.taskResponse(cctx, http.StatusCreated, task)
}

func (p *TaskPresenterImpl) UpdateTaskResponse(cctx fcontext.Context, task *entity.Task) error {
	return p.taskResponse(cctx, http.StatusOK, task)
}

func (p *TaskPresenterImpl) GetTaskResponse(cctx fcontext.Context, task *entity.Task) error {
	return p.taskResponse(cctx, http.StatusOK, task)
}

func (p *TaskPresenterImpl) taskResponse(cctx fcontext.Context, code int, task *entity.Task) error {
	ectx := fcontext.GetEchoContext(cctx)
	return ectx.JSON(code, TaskResponse(task))
}

func (p *TaskPresenterImpl) TaskAllResponse(cctx fcontext.Context, taskSlice entity.TaskSlice) error {
	ectx := fcontext.GetEchoContext(cctx)
	return ectx.JSON(http.StatusOK, GetAllTasksResponse(taskSlice))
}

func (p *TaskPresenterImpl) NilResponse(cctx fcontext.Context) error {
	ectx := fcontext.GetEchoContext(cctx)
	return ectx.NoContent(http.StatusOK)
}

func (p *TaskPresenterImpl) NoContentResponse(cctx fcontext.Context) error {
	ectx := fcontext.GetEchoContext(cctx)
	return ectx.NoContent(http.StatusNoContent)
}
