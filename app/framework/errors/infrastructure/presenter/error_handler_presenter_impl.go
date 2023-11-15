package presenter

import (
	"context"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	"net/http"

	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

type SystemErrorHandlerPresenterImpl struct {
}

func NewSystemErrorHandlerPresenterImpl() *SystemErrorHandlerPresenterImpl {
	return &SystemErrorHandlerPresenterImpl{}
}

func (p *SystemErrorHandlerPresenterImpl) ErrorResponse(ctx context.Context) error {
	ectx := fcontext.Ectx(ctx)
	message := "システムエラーが発生しました。"
	errorResponse := &openapi.ErrorResponse{Message: &message}
	return ectx.JSON(http.StatusInternalServerError, errorResponse)
}
