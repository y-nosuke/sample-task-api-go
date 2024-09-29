package presenter

import fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"

type BusinessErrorPresenter interface {
	BadRequest(ctx fcontext.Context, message string, err error) error
	Unauthorized(fcontext.Context, string) error
	Forbidden(ctx fcontext.Context, message string) error
	NotFound(ctx fcontext.Context, message string) error
	Conflict(ctx fcontext.Context, message string) error
}
