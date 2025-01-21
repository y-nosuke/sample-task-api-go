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
	registerTaskUseCase *usecase.CreateTaskUseCase
	getAllTaskUseCase   *usecase.GetAllTaskUseCase
	getTaskUseCase      *usecase.GetTaskUseCase
	updateTaskUseCase   *usecase.EditTaskUseCase
	deleteTaskUseCase   *usecase.DeleteTaskUseCase
	taskPresenter       presenter.TaskPresenter
}

func NewTaskHandler(registerTaskUseCase *usecase.CreateTaskUseCase,
	getAllTaskUseCase *usecase.GetAllTaskUseCase,
	getTaskUseCase *usecase.GetTaskUseCase,
	updateTaskUseCase *usecase.EditTaskUseCase,
	deleteTaskUseCase *usecase.DeleteTaskUseCase,
	taskPresenter presenter.TaskPresenter,
) *TaskHandler {
	return &TaskHandler{
		registerTaskUseCase,
		getAllTaskUseCase,
		getTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
		taskPresenter,
	}
}

func (h *TaskHandler) CreateTask(ectx echo.Context) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("create:task") {
		return ferrors.NewForbiddenError("指定された操作は許可されていません。 missing create:task")
	}

	request := new(openapi.CreateTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf("ectx.Bind(): %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		return ferrors.NewBadRequestErrorf(err, "バリデーションエラーです。")
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
		return ferrors.NewForbiddenError("指定された操作は許可されていません。 missing read:task")
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
		return ferrors.NewForbiddenError("指定された操作は許可されていません。 missing read:task")
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
		return ferrors.NewForbiddenError("指定された操作は許可されていません。 missing update:task")
	}

	request := new(openapi.EditTaskRequest)
	if err := ectx.Bind(request); err != nil {
		return xerrors.Errorf("ectx.Bind(): %w", err)
	}

	if err := ectx.Validate(request); err != nil {
		return ferrors.NewBadRequestErrorf(err, "バリデーションエラーです。")
	}

	args := EditTaskUseCaseArgs(id, request)
	if err := h.updateTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("updateTaskUseCase.Invoke(): %w", err)
	}

	return nil
}

func (h *TaskHandler) DeleteTask(ectx echo.Context, id uuid.UUID) error {
	ctx := context.CastContext(ectx)

	a := auth.GetAuth(ctx)
	if !a.HasAuthority("delete:task") {
		return ferrors.NewForbiddenError("指定された操作は許可されていません。 missing delete:task")
	}

	args := &usecase.DeleteTaskUseCaseArgs{Id: id}
	if err := h.deleteTaskUseCase.Invoke(ctx, args); err != nil {
		return xerrors.Errorf("deleteTaskUseCase.Invoke(): %w", err)
	}

	return nil
}
