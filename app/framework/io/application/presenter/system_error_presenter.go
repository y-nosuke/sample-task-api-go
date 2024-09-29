package presenter

import (
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
)

type SystemErrorPresenter interface {
	InternalServerError(ctx fcontext.Context) error
}
