package presenter

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator/v10"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/presenter/mapping"
	"golang.org/x/xerrors"
	"net/http"

	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

type TaskPresenterImpl struct {
}

func NewTaskPresenterImpl() *TaskPresenterImpl {
	return &TaskPresenterImpl{}
}

func (p *TaskPresenterImpl) RegisterTaskResponse(ctx context.Context, task *entity.Task) error {
	return p.taskResponse(ctx, http.StatusCreated, task)
}

func (p *TaskPresenterImpl) UpdateTaskResponse(ctx context.Context, task *entity.Task) error {
	return p.taskResponse(ctx, http.StatusOK, task)
}

func (p *TaskPresenterImpl) GetTaskResponse(ctx context.Context, task *entity.Task) error {
	return p.taskResponse(ctx, http.StatusOK, task)
}

func (p *TaskPresenterImpl) taskResponse(ctx context.Context, code int, task *entity.Task) error {
	ectx := fcontext.Ectx(ctx)
	taskForm := openapi.TaskResponse{
		TaskForm: *mapping.TaskForm(task),
	}

	return ectx.JSON(code, &taskForm)
}

func (p *TaskPresenterImpl) TaskAllResponse(ctx context.Context, tasks []*entity.Task) error {
	ectx := fcontext.Ectx(ctx)

	taskForm := openapi.GetAllTasksResponse{
		TaskForms: mapping.TaskForms(tasks),
	}

	return ectx.JSON(http.StatusOK, &taskForm)
}

func (p *TaskPresenterImpl) NilResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.NoContent(http.StatusOK)
}

func (p *TaskPresenterImpl) NoContentResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.NoContent(http.StatusNoContent)
}

func (p *TaskPresenterImpl) BadRequest(ctx context.Context, message string, err error) error {
	ectx := fcontext.Ectx(ctx)
	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		return xerrors.Errorf("errors.As()")
	}
	msg := message + "\n"
	msg += "errors: \n"
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println("=====================【validateチェック】===========================")
		fmt.Println("1: " + err.Namespace())
		fmt.Println("2: " + err.Field())
		fmt.Println("3: " + err.StructNamespace())
		fmt.Println("4: " + err.StructField())
		fmt.Println("5: " + err.Tag())
		fmt.Println("6: " + err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println("10" + err.Param())

		msg += "Namespace: " + err.Namespace() + "\n"
		msg += "Tag: " + err.Tag() + "\n"
		//message += "Value: " + err.Value().(string) + "\n"
	}
	fmt.Println("====================================================================")
	return ectx.JSON(http.StatusBadRequest, &openapi.ErrorResponse{Message: &msg})
}

func (p *TaskPresenterImpl) Forbidden(ctx context.Context, message string) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusForbidden, &openapi.ErrorResponse{Message: &message})
}

func (p *TaskPresenterImpl) NotFound(ctx context.Context, message string) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusNotFound, &openapi.ErrorResponse{Message: &message})
}

func (p *TaskPresenterImpl) Conflict(ctx context.Context, message string) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusConflict, &openapi.ErrorResponse{Message: &message})
}

func (p *TaskPresenterImpl) InternalServerError(ctx context.Context, message string) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusInternalServerError, &openapi.ErrorResponse{Message: &message})
}
