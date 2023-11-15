package presenter

import (
	"context"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"net/http"
)

type AuthHandlerPresenterImpl struct {
}

func NewAuthHandlerPresenterImpl() *AuthHandlerPresenterImpl {
	return &AuthHandlerPresenterImpl{}
}

func (p *AuthHandlerPresenterImpl) Unauthorized(ctx context.Context, message string) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusUnauthorized, &openapi.ErrorResponse{Message: &message})
}
