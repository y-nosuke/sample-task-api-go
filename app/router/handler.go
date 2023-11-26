package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func customHTTPErrorHandler(err error, ectx echo.Context) {
	ectx.Logger().Error(err)
	fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
}
