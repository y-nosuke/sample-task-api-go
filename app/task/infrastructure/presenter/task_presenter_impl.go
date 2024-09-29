package presenter

import (
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
