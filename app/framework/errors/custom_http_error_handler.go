package errors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func CustomHTTPErrorHandler(err error, ectx echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		ectx.Logger().Warnf("%+v\n", xerrors.Errorf("http error!: %w", err))
	} else {
		ectx.Logger().Error("%+v\n", xerrors.Errorf("system error!: %w", err))
	}
}
