package presenter

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"net/http"

	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
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
		case ferrors.BadRequest:
			return http.StatusBadRequest
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
		switch appError.Status {
		case ferrors.BadRequest:
			if _, ok := appError.OriginalError.(*validator.InvalidValidationError); ok {
				fmt.Println(appError.OriginalError)
				message := "システムエラーが発生しました。"
				return &openapi.ErrorResponse{Message: &message}
			}
			message := appError.Message + "\n"
			message += "errors: \n"
			for _, err := range appError.OriginalError.(validator.ValidationErrors) {
				fmt.Println("=====================【validateチェック】===========================")
				fmt.Println("1: " + err.Namespace())
				fmt.Println("2: " + err.Field())
				fmt.Println("3: " + err.StructNamespace())
				fmt.Println("4: " + err.StructField())
				fmt.Println("5: " + err.Tag())
				fmt.Println("6: " + err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println("10" + err.Param())

				message += "Namespace: " + err.Namespace() + "\n"
				message += "Tag: " + err.Tag() + "\n"
				//message += "Value: " + err.Value().(string) + "\n"
			}
			fmt.Println("====================================================================")

			return &openapi.ErrorResponse{Message: &message}
		default:
			return &openapi.ErrorResponse{Message: &appError.Message}
		}
	} else if errors.As(err, &httpError) {
		message := httpError.Message.(string)
		return &openapi.ErrorResponse{Message: &message}
	} else {
		message := "システムエラーが発生しました。"
		return &openapi.ErrorResponse{Message: &message}
	}
}
