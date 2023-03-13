package controllers

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
	"github.com/y-nosuke/sample-task-api-go/task/application/usecases"
	"golang.org/x/xerrors"
)

type TaskController struct {
	registerTaskUseCase *usecases.RegisterTaskUseCase
}

func NewTaskController(registerTaskUseCase *usecases.RegisterTaskUseCase) *TaskController {
	return &TaskController{registerTaskUseCase}
}

func (c *TaskController) RegisterTask(ectx echo.Context) error {
	fmt.Println("タスクの登録処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	request := new(openapi.RegisterTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	fmt.Println("request: ", request)

	args := request2Args(request)

	if err := c.registerTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func request2Args(request *openapi.RegisterTaskRequest) *usecases.RegisterTaskUseCaseArgs {
	var detail string
	if request.Detail != nil {
		detail = *request.Detail
	}

	var deadline time.Time
	if request.Deadline != nil {
		deadline = request.Deadline.Time
	}

	return &usecases.RegisterTaskUseCaseArgs{
		Title:    request.Title,
		Detail:   detail,
		Deadline: deadline,
	}
}
