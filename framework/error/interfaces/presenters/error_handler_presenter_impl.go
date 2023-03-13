package presenters

import (
	"context"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
)

type ErrorHandlerPresenter struct {
}

func NewErrorHandlerPresenter() *ErrorHandlerPresenter {
	return &ErrorHandlerPresenter{}
}

func (p *ErrorHandlerPresenter) ErrorResponse(ctx context.Context, err error) error {
	ectx := fcontext.Ectx(ctx)
	return ectx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
}
