package handler

import (
	"fmt"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
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
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing create:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	request := new(openapi.RegisterTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return errors.SystemErrorf("ectx.Bind(): %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		if err := h.taskPresenter.BadRequest(ctx, "バリデーションエラーです。", err); err != nil {
			return errors.SystemErrorf("taskPresenter.BadRequest(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.BadRequest()")
	}

	args := RegisterTaskUseCaseArgs(request)

	if err := h.registerTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("registerTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) GetAllTasks(ectx echo.Context) error {
	fmt.Println("タスク一覧取得処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("read:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing read:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	args := &usecase.GetAllTaskUseCaseArgs{}

	if err := h.getAllTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("getAllTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) GetTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク取得処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("read:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing read:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	args := &usecase.GetTaskUseCaseArgs{Id: id}

	if err := h.getTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("getTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) UpdateTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク更新処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	request := new(openapi.UpdateTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return errors.SystemErrorf("ectx.Bind(): %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		if err := h.taskPresenter.BadRequest(ctx, "バリデーションエラーです。", err); err != nil {
			return errors.SystemErrorf("taskPresenter.BadRequest(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.BadRequest()")
	}

	args := UpdateTaskUseCaseArgs(id, request)

	if err := h.updateTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("updateTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) CompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク完了処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	request := new(openapi.CompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return errors.SystemErrorf("ectx.Bind(): %w", err)
	}

	args := CompleteTaskUseCaseArgs(id, request)

	if err := h.completeTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("completeTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) UnCompleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク未完了処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	request := new(openapi.UnCompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return errors.SystemErrorf("ectx.Bind(): %w", err)
	}

	args := UnCompleteTaskUseCaseArgs(id, request)

	if err := h.unCompleteTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("unCompleteTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) DeleteTask(ectx echo.Context, id uuid.UUID) error {
	fmt.Println("タスク削除処理を開始します。")
	ctx := context.Ctx(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("delete:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing delete:task"); err != nil {
			return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
		}
		return errors.BusinessErrorf("taskPresenter.Forbidden()")
	}

	args := &usecase.DeleteTaskUseCaseArgs{Id: id}

	if err := h.deleteTaskUseCase.Invoke(ctx, args); err != nil {
		return errors.SystemErrorf("deleteTaskUseCase.Invoke(): %w", err)
	}

	return nil
}
