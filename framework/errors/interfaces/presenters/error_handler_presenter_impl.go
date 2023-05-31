package presenters

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"net/http"

	ferrors "github.com/y-nosuke/sample-task-api-go/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
)

type ErrorHandlerPresenter struct {
}

func NewErrorHandlerPresenter() *ErrorHandlerPresenter {
	return &ErrorHandlerPresenter{}
}

func (p *ErrorHandlerPresenter) ErrorResponse(ctx context.Context, err error) error {
	fmt.Println(err.Error())

	ectx := fcontext.Ectx(ctx)
	httpStatus := httpStatus(err)
	errorResponse := errorResponse(err)

	return ectx.JSON(httpStatus, errorResponse)
}

func httpStatus(err error) int {
	var appError *ferrors.AppError
	var httpError *echo.HTTPError
	if errors.As(err, &appError) {
		switch appError.Status {
		case ferrors.Unauthorized:
			return http.StatusUnauthorized
		case ferrors.Forbidden:
			return http.StatusForbidden
		case ferrors.NotFound:
			return http.StatusNotFound
		case ferrors.Conflict:
			return http.StatusConflict
		default:
			return http.StatusInternalServerError
		}
	} else if errors.As(err, &httpError) {
		return httpError.Code
	} else {
		return http.StatusInternalServerError
	}
}

func errorResponse(err error) *openapi.ErrorResponse {
	var appError *ferrors.AppError
	var httpError *echo.HTTPError
	if errors.As(err, &appError) {
		return &openapi.ErrorResponse{Message: &appError.Message}
	} else if errors.As(err, &httpError) {
		message := httpError.Message.(string)
		return &openapi.ErrorResponse{Message: &message}
	} else {
		message := "システムエラーが発生しました。"
		return &openapi.ErrorResponse{Message: &message}
	}
}
