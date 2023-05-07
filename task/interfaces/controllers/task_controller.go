package controllers

import (
	"fmt"
	auth2 "github.com/y-nosuke/sample-task-api-go/framework/auth"
	fauth "github.com/y-nosuke/sample-task-api-go/framework/auth/interfaces"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	ferrors "github.com/y-nosuke/sample-task-api-go/framework/errors"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
	"github.com/y-nosuke/sample-task-api-go/task/application/usecases"
	"golang.org/x/xerrors"
)

type TaskController struct {
	registerTaskUseCase   *usecases.RegisterTaskUseCase
	getAllTaskUseCase     *usecases.GetAllTaskUseCase
	getTaskUseCase        *usecases.GetTaskUseCase
	updateTaskUseCase     *usecases.UpdateTaskUseCase
	completeTaskUseCase   *usecases.CompleteTaskUseCase
	unCompleteTaskUseCase *usecases.UnCompleteTaskUseCase
	deleteTaskUseCase     *usecases.DeleteTaskUseCase
}

func NewTaskController(registerTaskUseCase *usecases.RegisterTaskUseCase,
	getAllTaskUseCase *usecases.GetAllTaskUseCase,
	getTaskUseCase *usecases.GetTaskUseCase,
	updateTaskUseCase *usecases.UpdateTaskUseCase,
	completeTaskUseCase *usecases.CompleteTaskUseCase,
	unCompleteTaskUseCase *usecases.UnCompleteTaskUseCase,
	deleteTaskUseCase *usecases.DeleteTaskUseCase,
) *TaskController {
	return &TaskController{
		registerTaskUseCase,
		getAllTaskUseCase,
		getTaskUseCase,
		updateTaskUseCase,
		completeTaskUseCase,
		unCompleteTaskUseCase,
		deleteTaskUseCase,
	}
}

func (c *TaskController) RegisterTask(ectx echo.Context) error {
	fmt.Println("タスク登録処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("create:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing create:task"))
	}

	request := new(openapi.RegisterTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	fmt.Println("request: ", request)

	args := request2RegisterArgs(request)

	if err := c.registerTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (c *TaskController) GetAllTasks(ectx echo.Context) error {
	fmt.Println("タスク一覧取得処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("read:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing read:task"))
	}

	args := &usecases.GetAllTaskUseCaseArgs{}

	if err := c.getAllTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (c *TaskController) GetTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク取得処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("read:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing read:task"))
	}

	args := &usecases.GetTaskUseCaseArgs{Id: id}

	if err := c.getTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (c *TaskController) UpdateTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク更新処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("update:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing update:task"))
	}

	request := new(openapi.UpdateTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	fmt.Println("request: ", request)

	args, err := request2UpdateArgs(id, request)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := c.updateTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (c *TaskController) CompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク完了処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("update:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing update:task"))
	}

	request := new(openapi.CompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	fmt.Println("request: ", request)

	args, err := request2CompleteArgs(id, request)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := c.completeTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (c *TaskController) UnCompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク未完了処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("update:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing update:task"))
	}

	request := new(openapi.UnCompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	fmt.Println("request: ", request)

	args, err := request2UnCompleteArgs(id, request)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := c.unCompleteTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (c *TaskController) DeleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク削除処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	auth := cctx.Value(fauth.AUTH).(*auth2.Auth)
	if !auth.HasAuthority("delete:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing delete:task"))
	}

	args := &usecases.DeleteTaskUseCaseArgs{Id: id}

	if err := c.deleteTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func request2RegisterArgs(request *openapi.RegisterTaskRequest) *usecases.RegisterTaskUseCaseArgs {
	var deadline *time.Time
	if request.Deadline != nil {
		deadline = &request.Deadline.Time
	}

	return &usecases.RegisterTaskUseCaseArgs{
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: deadline,
	}
}

func request2UpdateArgs(id uuid.UUID, request *openapi.UpdateTaskRequest) (*usecases.UpdateTaskUseCaseArgs, error) {
	return &usecases.UpdateTaskUseCaseArgs{
		Id:        id,
		Title:     request.Title,
		Detail:    request.Detail,
		Completed: *request.Completed,
		Deadline:  &request.Deadline.Time,
		Version:   request.Version,
	}, nil
}

func request2CompleteArgs(id uuid.UUID, request *openapi.CompleteTaskRequest) (*usecases.CompleteTaskUseCaseArgs, error) {
	return &usecases.CompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}

func request2UnCompleteArgs(id uuid.UUID, request *openapi.UnCompleteTaskRequest) (*usecases.UnCompleteTaskUseCaseArgs, error) {
	return &usecases.UnCompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}
