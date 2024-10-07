package errors

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, ectx echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		ectx.Logger().Infof("http error!: %+v", he)
		fmt.Printf("http error!: %+v\n", he)
	} else if businessError, ok := AsBusinessError(err); ok {
		ectx.Logger().Warnf("business error!: %+v, orinal error: %+v", businessError, businessError.OriginalError())
		fmt.Printf("business error!: %+v, orinal error: %+v\n", businessError, businessError.OriginalError())
	} else {
		ectx.Logger().Errorf("system error!: %+v", err)
		fmt.Printf("system error!: %+v\n", err)
	}
}
