package presenter

import (
	"context"
)

type ErrorHandlerPresenter interface {
	ErrorResponse(context.Context, error) error
}
