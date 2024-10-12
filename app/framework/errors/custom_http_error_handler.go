package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, ectx echo.Context) {
	if ectx.Response().Committed {
		// middlewareでecho.Context.Error(err error)を呼ばれるとHTTPErrorHandlerが呼ばれる可能性があるので、この判定を行う
		return
	}

	if resErr := ectx.NoContent(http.StatusInternalServerError); resErr != nil {
		panic(fmt.Sprintf("error occurred. %v", resErr))
	}
}
