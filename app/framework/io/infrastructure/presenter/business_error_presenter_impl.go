package presenter

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

type BusinessErrorPresenterImpl struct {
}

func NewBusinessErrorPresenterImpl() *BusinessErrorPresenterImpl {
	return &BusinessErrorPresenterImpl{}
}

func (p BusinessErrorPresenterImpl) BadRequest(ctx context.Context, message string, err error) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusBadRequest, &openapi.ErrorResponse{Message: BadRequestMessage(message, err)})
}

func (p BusinessErrorPresenterImpl) Unauthorized(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusUnauthorized, &openapi.ErrorResponse{Message: message})
}

func (p BusinessErrorPresenterImpl) Forbidden(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusForbidden, &openapi.ErrorResponse{Message: message})
}

func (p BusinessErrorPresenterImpl) NotFound(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusNotFound, &openapi.ErrorResponse{Message: message})
}

func (p BusinessErrorPresenterImpl) Conflict(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusConflict, &openapi.ErrorResponse{Message: message})
}
