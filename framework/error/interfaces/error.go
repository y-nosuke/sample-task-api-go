package interfaces

import (
	"fmt"

	"github.com/labstack/echo/v4"
	ferror "github.com/y-nosuke/sample-task-api-go/framework/error/interfaces/presenters"
	"golang.org/x/xerrors"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("エラーハンドラーを実行します。")

		if err := next(ectx); err != nil {
			fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))

			errorHandlerPresenter := ferror.NewErrorHandlerPresenter()
			return errorHandlerPresenter.ErrorResponse(ectx, err)
		}

		return nil
	}
}
