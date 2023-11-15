package presenter

import "context"

// TODO: interfaceが必要か考える
type AuthHandlerPresenter interface {
	Unauthorized(context.Context, string) error
}
