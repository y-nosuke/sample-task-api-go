package presenter

import (
	"context"
)

// TODO: interfaceが必要か考える
type SystemErrorHandlerPresenter interface {
	ErrorResponse(context.Context) error
}
