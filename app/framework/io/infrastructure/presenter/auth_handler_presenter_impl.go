package presenter

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

type AuthHandlerPresenterImpl struct {
}

func NewAuthHandlerPresenterImpl() *AuthHandlerPresenterImpl {
	return &AuthHandlerPresenterImpl{}
}

func (p *AuthHandlerPresenterImpl) Unauthorized(ctx context.Context, message string) error {
	ectx := fcontext.GetEctx(ctx)
	return ectx.JSON(http.StatusUnauthorized, &openapi.ErrorResponse{Message: message})
}
