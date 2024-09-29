package errors

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func CustomHTTPErrorHandler(err error, ectx echo.Context) {
	ectx.Logger().Error(err)
	fmt.Printf("%+v\n", xerrors.Errorf("system error!: %w", err))
}
