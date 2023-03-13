package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
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

	args := new(usecases.RegisterTaskUseCaseArgs)
	if err := ectx.Bind(args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := c.registerTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
