package interfaces

import (
	"fmt"
	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors/infrastructure/presenter"
	"golang.org/x/xerrors"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("エラーハンドラーを実行します。")
		cctx := fcontext.Cctx(ectx)

		if err := next(ectx); err != nil {
			fmt.Println("エラーハンドラー")
			fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))

			errorHandlerPresenterImpl := ferrors.NewErrorHandlerPresenterImpl()
			return errorHandlerPresenterImpl.ErrorResponse(cctx.Ctx, err)
		}

		return nil
	}
}
