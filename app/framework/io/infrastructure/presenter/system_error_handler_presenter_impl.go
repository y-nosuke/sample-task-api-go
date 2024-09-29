package presenter

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"

	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

type SystemErrorHandlerPresenterImpl struct {
}

func NewSystemErrorHandlerPresenterImpl() *SystemErrorHandlerPresenterImpl {
	return &SystemErrorHandlerPresenterImpl{}
}

func (p *SystemErrorHandlerPresenterImpl) ErrorResponse(ctx context.Context) error {
	ectx := fcontext.GetEctx(ctx)
	errorResponse := &openapi.ErrorResponse{Message: "システムエラーが発生しました。"}
	return ectx.JSON(http.StatusInternalServerError, errorResponse)
}
