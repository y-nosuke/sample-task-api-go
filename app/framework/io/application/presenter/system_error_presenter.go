package presenter

import (
	"context"
)

type SystemErrorPresenter interface {
	InternalServerError(ctx context.Context) error
}
