package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"github.com/y-nosuke/sample-task-api-go/task/application/usecases"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/presenters"
	"golang.org/x/xerrors"
)

type TaskController struct {
	registerTaskUseCase *usecases.RegisterTaskUseCase
	taskPresenter       *presenters.TaskPresenter
}

func NewTaskController(registerTaskUseCase *usecases.RegisterTaskUseCase, taskPresenter *presenters.TaskPresenter) *TaskController {
	return &TaskController{registerTaskUseCase, taskPresenter}
}

func (c *TaskController) RegisterTask(ectx echo.Context) error {
	fmt.Println("タスクの登録処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	inputData := new(usecases.RegisterTaskUseCaseInputData)
	if err := ectx.Bind(inputData); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	outputData, err := c.registerTaskUseCase.Invoke(cctx.Ctx, inputData)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := c.taskPresenter.TaskResponse(ectx, outputData.Task); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
