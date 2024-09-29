package presenter

import (
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"net/http"
)

type SystemErrorHandlerPresenterImpl struct {
}

func NewSystemErrorHandlerPresenterImpl() *SystemErrorHandlerPresenterImpl {
	return &SystemErrorHandlerPresenterImpl{}
}

func (p SystemErrorHandlerPresenterImpl) InternalServerError(ctx fcontext.Context) error {
	ectx := fcontext.GetEchoContext(ctx)
	return ectx.JSON(http.StatusInternalServerError, &openapi.ErrorResponse{Message: "システムエラーが発生しました。"})
}
