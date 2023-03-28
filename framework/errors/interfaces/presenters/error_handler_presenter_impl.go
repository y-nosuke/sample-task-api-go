package presenters

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"net/http"

	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
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
	errorResponse := appError2ErrorResponse(err)

	return ectx.JSON(httpStatus, errorResponse)
}

func httpStatus(err error) int {
	var appError *ferrors.AppError
	if errors.As(err, &appError) {
		fmt.Println(1)
		fmt.Printf("%T", appError)
		switch appError.Status {
		case ferrors.NotFound:
			return http.StatusNotFound
		case ferrors.Conflict:
			return http.StatusConflict
		default:
			return http.StatusInternalServerError
		}
	} else {
		fmt.Println(4)
		fmt.Printf("%T", err)
		return http.StatusInternalServerError
	}
}

func appError2ErrorResponse(err error) *openapi.ErrorResponse {
	var appError *ferrors.AppError
	if errors.As(err, &appError) {
		return &openapi.ErrorResponse{Message: &appError.Message}
	} else {
		message := "システムエラーが発生しました。"
		return &openapi.ErrorResponse{Message: &message}
	}
}
