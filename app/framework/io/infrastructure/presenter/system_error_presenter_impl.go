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

func (p SystemErrorHandlerPresenterImpl) InternalServerError(ctx context.Context) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusInternalServerError, &openapi.ErrorResponse{Message: "システムエラーが発生しました。"})
}
