package interfaces

import (
	"fmt"

	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	ferror "github.com/y-nosuke/sample-task-api-go/framework/error/interfaces/presenters"
	"golang.org/x/xerrors"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("エラーハンドラーを実行します。")
		cctx := fcontext.Cctx(ectx)

		if err := next(ectx); err != nil {
			fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))

			errorHandlerPresenter := ferror.NewErrorHandlerPresenter()
			return errorHandlerPresenter.ErrorResponse(cctx.Ctx, err)
		}

		return nil
	}
}
