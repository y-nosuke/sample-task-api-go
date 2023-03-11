package presenters

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorHandlerPresenter struct {
}

func NewErrorHandlerPresenter() *ErrorHandlerPresenter {
	return &ErrorHandlerPresenter{}
}

func (p *ErrorHandlerPresenter) ErrorResponse(ectx echo.Context, err error) error {
	return ectx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
}
