package handler

import (
	"fmt"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"golang.org/x/xerrors"
)

type TaskHandler struct {
	registerTaskUseCase   *usecase.RegisterTaskUseCase
	getAllTaskUseCase     *usecase.GetAllTaskUseCase
	getTaskUseCase        *usecase.GetTaskUseCase
	updateTaskUseCase     *usecase.UpdateTaskUseCase
	completeTaskUseCase   *usecase.CompleteTaskUseCase
	unCompleteTaskUseCase *usecase.UnCompleteTaskUseCase
	deleteTaskUseCase     *usecase.DeleteTaskUseCase
}

func NewTaskHandler(registerTaskUseCase *usecase.RegisterTaskUseCase,
	getAllTaskUseCase *usecase.GetAllTaskUseCase,
	getTaskUseCase *usecase.GetTaskUseCase,
	updateTaskUseCase *usecase.UpdateTaskUseCase,
	completeTaskUseCase *usecase.CompleteTaskUseCase,
	unCompleteTaskUseCase *usecase.UnCompleteTaskUseCase,
	deleteTaskUseCase *usecase.DeleteTaskUseCase,
) *TaskHandler {
	return &TaskHandler{
		registerTaskUseCase,
		getAllTaskUseCase,
		getTaskUseCase,
		updateTaskUseCase,
		completeTaskUseCase,
		unCompleteTaskUseCase,
		deleteTaskUseCase,
	}
}

func (h *TaskHandler) RegisterTask(ectx echo.Context) error {
	fmt.Println("タスク登録処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("create:task") {
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

	if err := h.registerTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) GetAllTasks(ectx echo.Context) error {
	fmt.Println("タスク一覧取得処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("read:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing read:task"))
	}

	args := &usecase.GetAllTaskUseCaseArgs{}

	if err := h.getAllTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) GetTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク取得処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("read:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing read:task"))
	}

	args := &usecase.GetTaskUseCaseArgs{Id: id}

	if err := h.getTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) UpdateTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク更新処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("update:task") {
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

	if err := h.updateTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) CompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク完了処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("update:task") {
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

	if err := h.completeTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) UnCompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク未完了処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("update:task") {
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

	if err := h.unCompleteTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) DeleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク削除処理を開始します。")
	cctx := fcontext.Cctx(ectx)

	a := cctx.Value(fauth.AUTH).(*auth.Auth)
	if !a.HasAuthority("delete:task") {
		return ferrors.New(ferrors.Forbidden, "指定された操作は許可されていません。", fmt.Errorf("missing delete:task"))
	}

	args := &usecase.DeleteTaskUseCaseArgs{Id: id}

	if err := h.deleteTaskUseCase.Invoke(cctx.Ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func registerTaskUseCaseArgs(request *openapi.RegisterTaskRequest) *usecase.RegisterTaskUseCaseArgs {
	var deadline *time.Time
	if request.Deadline != nil {
		deadline = &request.Deadline.Time
	}

	return &usecase.RegisterTaskUseCaseArgs{
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: deadline,
	}
}

func updateTaskUseCaseArgs(id uuid.UUID, request *openapi.UpdateTaskRequest) (*usecase.UpdateTaskUseCaseArgs, error) {
	return &usecase.UpdateTaskUseCaseArgs{
		Id:       id,
		Title:    request.Title,
		Detail:   request.Detail,
		Deadline: &request.Deadline.Time,
		Version:  &request.Version,
	}, nil
}

func completeTaskUseCaseArgs(id uuid.UUID, request *openapi.CompleteTaskRequest) (*usecase.CompleteTaskUseCaseArgs, error) {
	return &usecase.CompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}

func unCompleteTaskUseCaseArgs(id uuid.UUID, request *openapi.UnCompleteTaskRequest) (*usecase.UnCompleteTaskUseCaseArgs, error) {
	return &usecase.UnCompleteTaskUseCaseArgs{
		Id:      id,
		Version: &request.Version,
	}, nil
}
