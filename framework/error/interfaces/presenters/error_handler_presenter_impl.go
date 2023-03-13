package presenters

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
)

type ErrorHandlerPresenter struct {
}

func NewErrorHandlerPresenter() *ErrorHandlerPresenter {
	return &ErrorHandlerPresenter{}
}

func (p *ErrorHandlerPresenter) ErrorResponse(ctx context.Context, err error) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusInternalServerError, error2ErrorResponse(err))
}

func error2ErrorResponse(err error) *openapi.ErrorResponse {
	message := err.Error()
	return &openapi.ErrorResponse{Message: &message}
}
