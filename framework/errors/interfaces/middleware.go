package interfaces

import (
	"fmt"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"

	"github.com/labstack/echo/v4"
	ferrors "github.com/y-nosuke/sample-task-api-go/framework/errors/interfaces/presenters"
	"golang.org/x/xerrors"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("エラーハンドラーを実行します。")
		cctx := fcontext.Cctx(ectx)

		if err := next(ectx); err != nil {
			fmt.Println("エラーハンドラー")
			fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))

			errorHandlerPresenter := ferrors.NewErrorHandlerPresenter()
			return errorHandlerPresenter.ErrorResponse(cctx.Ctx, err)
		}

		return nil
	}
}
