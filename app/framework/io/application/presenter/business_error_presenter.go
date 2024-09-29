package presenter

import "context"

type BusinessErrorPresenter interface {
	BadRequest(ctx context.Context, message string, err error) error
	Unauthorized(context.Context, string) error
	Forbidden(ctx context.Context, message string) error
	NotFound(ctx context.Context, message string) error
	Conflict(ctx context.Context, message string) error
}
