package presenters

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
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
	} else {
		return http.StatusInternalServerError
	}
}

func errorResponse(err error) *openapi.ErrorResponse {
	var appError *ferrors.AppError
	if errors.As(err, &appError) {
		return &openapi.ErrorResponse{Message: &appError.Message}
	} else {
		message := "システムエラーが発生しました。"
		return &openapi.ErrorResponse{Message: &message}
	}
}
