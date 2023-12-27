package interfaces

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"golang.org/x/xerrors"
)

func CustomHTTPErrorHandler(err error, ectx echo.Context) {
	var businessError *ferrors.BusinessError
	if errors.Is(err, businessError) {
		ectx.Logger().Warn(err)
		fmt.Printf("%+v\n", err)
		return
	}

	ectx.Logger().Error(err)
	fmt.Printf("%+v\n", xerrors.Errorf("system error!: %w", err))
}
