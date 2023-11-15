package controller

import (
	"fmt"
	auth2 "github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	usecase2 "github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"golang.org/x/xerrors"
)

type TaskController struct {
	registerTaskUseCase   *usecase2.RegisterTaskUseCase
	getAllTaskUseCase     *usecase2.GetAllTaskUseCase
	getTaskUseCase        *usecase2.GetTaskUseCase
	updateTaskUseCase     *usecase2.UpdateTaskUseCase
	completeTaskUseCase   *usecase2.CompleteTaskUseCase
	unCompleteTaskUseCase *usecase2.UnCompleteTaskUseCase
	deleteTaskUseCase     *usecase2.DeleteTaskUseCase
}

func NewTaskController(registerTaskUseCase *usecase2.RegisterTaskUseCase,
	getAllTaskUseCase *usecase2.GetAllTaskUseCase,
	getTaskUseCase *usecase2.GetTaskUseCase,
	updateTaskUseCase *usecase2.UpdateTaskUseCase,
	completeTaskUseCase *usecase2.CompleteTaskUseCase,
	unCompleteTaskUseCase *usecase2.UnCompleteTaskUseCase,
	deleteTaskUseCase *usecase2.DeleteTaskUseCase,
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

	if err := ectx.Validate(request); err != nil {
		return ferrors.New(ferrors.BadRequest, "バリデーションエラーです。", err)
	}

	args := registerTaskUseCaseArgs(request)

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

	args := &usecase2.GetAllTaskUseCaseArgs{}

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

	args := &usecase2.GetTaskUseCaseArgs{Id: id}

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

	if err := ectx.Validate(request); err != nil {
		return ferrors.New(ferrors.BadRequest, "バリデーションエラーです。", err)
	}

	args, err := updateTaskUseCaseArgs(id, request)
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

	args, err := completeTaskUseCaseArgs(id, request)
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

	args, err := unCompleteTaskUseCaseArgs(id, request)
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

	args := &usecase2.DeleteTaskUseCaseArgs{Id: id}

	if err := c.deleteTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func registerTaskUseCaseArgs(request *openapi.RegisterTaskRequest) *usecase2.RegisterTaskUseCaseArgs {
	var deadline *time.Time
	if request.Deadline != nil {
		deadline = &request.Deadline.Time
	}

	return &usecase2.RegisterTaskUseCaseArgs{
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: deadline,
	}
}

func updateTaskUseCaseArgs(id uuid.UUID, request *openapi.UpdateTaskRequest) (*usecase2.UpdateTaskUseCaseArgs, error) {
	return &usecase2.UpdateTaskUseCaseArgs{
		Id:       id,
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: &request.Deadline.Time,
		Version:  &request.Version,
	}, nil
}

func completeTaskUseCaseArgs(id uuid.UUID, request *openapi.CompleteTaskRequest) (*usecase2.CompleteTaskUseCaseArgs, error) {
	return &usecase2.CompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}

func unCompleteTaskUseCaseArgs(id uuid.UUID, request *openapi.UnCompleteTaskRequest) (*usecase2.UnCompleteTaskUseCaseArgs, error) {
	return &usecase2.UnCompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}
