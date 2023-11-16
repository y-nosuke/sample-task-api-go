package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/handler/mapping"
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
	taskPresenter         presenter.TaskPresenter
}

func NewTaskHandler(registerTaskUseCase *usecase.RegisterTaskUseCase,
	getAllTaskUseCase *usecase.GetAllTaskUseCase,
	getTaskUseCase *usecase.GetTaskUseCase,
	updateTaskUseCase *usecase.UpdateTaskUseCase,
	completeTaskUseCase *usecase.CompleteTaskUseCase,
	unCompleteTaskUseCase *usecase.UnCompleteTaskUseCase,
	deleteTaskUseCase *usecase.DeleteTaskUseCase,
	taskPresenter presenter.TaskPresenter,
) *TaskHandler {
	return &TaskHandler{
		registerTaskUseCase,
		getAllTaskUseCase,
		getTaskUseCase,
		updateTaskUseCase,
		completeTaskUseCase,
		unCompleteTaskUseCase,
		deleteTaskUseCase,
		taskPresenter,
	}
}

func (h *TaskHandler) RegisterTask(ectx echo.Context) error {
	fmt.Println("タスク登録処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("create:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing create:task")
	}

	request := new(openapi.RegisterTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		return h.taskPresenter.BadRequest(ctx, "バリデーションエラーです。", err)
	}

	args := mapping.RegisterTaskUseCaseArgs(request)

	if err := h.registerTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) GetAllTasks(ectx echo.Context) error {
	fmt.Println("タスク一覧取得処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("read:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing read:task")
	}

	args := &usecase.GetAllTaskUseCaseArgs{}

	if err := h.getAllTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) GetTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク取得処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("read:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing read:task")
	}

	args := &usecase.GetTaskUseCaseArgs{Id: id}

	if err := h.getTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) UpdateTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク更新処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.UpdateTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		return h.taskPresenter.BadRequest(ctx, "バリデーションエラーです。", err)
	}

	args, err := mapping.UpdateTaskUseCaseArgs(id, request)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := h.updateTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) CompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク完了処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.CompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	args, err := mapping.CompleteTaskUseCaseArgs(id, request)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := h.completeTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) UnCompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク未完了処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.UnCompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	args, err := mapping.UnCompleteTaskUseCaseArgs(id, request)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := h.unCompleteTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (h *TaskHandler) DeleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク削除処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("delete:task") {
		return h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing delete:task")
	}

	args := &usecase.DeleteTaskUseCaseArgs{Id: id}

	if err := h.deleteTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
