package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"golang.org/x/xerrors"
)

type TaskHandler struct {
	registerTaskUseCase   *usecase.CreateTaskUseCase
	getAllTaskUseCase     *usecase.GetAllTaskUseCase
	getTaskUseCase        *usecase.GetTaskUseCase
	updateTaskUseCase     *usecase.EditTaskUseCase
	completeTaskUseCase   *usecase.CompleteTaskUseCase
	unCompleteTaskUseCase *usecase.UnCompleteTaskUseCase
	deleteTaskUseCase     *usecase.DeleteTaskUseCase
	taskPresenter         presenter.TaskPresenter
}

func NewTaskHandler(registerTaskUseCase *usecase.CreateTaskUseCase,
	getAllTaskUseCase *usecase.GetAllTaskUseCase,
	getTaskUseCase *usecase.GetTaskUseCase,
	updateTaskUseCase *usecase.EditTaskUseCase,
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

func (h *TaskHandler) CreateTask(ectx echo.Context) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("create:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing create:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing create:task")
	}

	request := new(openapi.CreateTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf("ectx.Bind(): %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		if resErr := h.taskPresenter.BadRequest(ctx, "バリデーションエラーです。", err); resErr != nil {
			return xerrors.Errorf("original err: %+v, taskPresenter.BadRequest(): %w", err, resErr)
		}
		return ferrors.NewBusinessError("バリデーションエラーです。")
	}

	args := RegisterTaskUseCaseArgs(request)
	if err := h.registerTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("registerTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) GetAllTasks(ectx echo.Context) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("read:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing read:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing read:task")
	}

	args := &usecase.GetAllTaskUseCaseArgs{}
	if err := h.getAllTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("getAllTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) GetTask(ectx echo.Context, id uuid.UUID) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("read:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing read:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing read:task")
	}

	args := &usecase.GetTaskUseCaseArgs{Id: id}
	if err := h.getTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("getTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) EditTask(ectx echo.Context, id uuid.UUID) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.EditTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf("ectx.Bind(): %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		if resErr := h.taskPresenter.BadRequest(ctx, "バリデーションエラーです。", err); resErr != nil {
			return xerrors.Errorf("original err: %+v, taskPresenter.BadRequest(): %w", err, resErr)
		}
		return ferrors.NewBusinessError("バリデーションエラーです。")
	}

	args := UpdateTaskUseCaseArgs(id, request)
	if err := h.updateTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("updateTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) CompleteTask(ectx echo.Context, id uuid.UUID) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.CompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf("ectx.Bind(): %w", err)
	}

	args := CompleteTaskUseCaseArgs(id, request)
	if err := h.completeTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("completeTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) UnCompleteTask(ectx echo.Context, id uuid.UUID) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("update:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing update:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.UnCompleteTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf("ectx.Bind(): %w", err)
	}

	args := UnCompleteTaskUseCaseArgs(id, request)
	if err := h.unCompleteTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("unCompleteTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) DeleteTask(ectx echo.Context, id uuid.UUID) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("delete:task") {
		if err := h.taskPresenter.Forbidden(ctx, "指定された操作は許可されていません。 missing delete:task"); err != nil {
			return xerrors.Errorf("taskPresenter.Forbidden(): %w", err)
		}
		return ferrors.NewBusinessError("指定された操作は許可されていません。 missing delete:task")
	}

	args := &usecase.DeleteTaskUseCaseArgs{Id: id}
	if err := h.deleteTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("deleteTaskUseCase.Invoke(): %w", err)
	}

	return nil
}
