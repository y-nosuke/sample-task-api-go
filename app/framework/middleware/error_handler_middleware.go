package middleware

import (
	"fmt"
	"github.com/y-nosuke/sample-task-api-go/app/framework/io/application/presenter"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"golang.org/x/xerrors"
)

func ErrorHandlerMiddleware(systemErrorHandlerPresenter presenter.SystemErrorHandlerPresenter) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			fmt.Println("エラーハンドラーを実行します。")
			cctx := context.Cctx(ectx)

			if err := next(ectx); err != nil {
				fmt.Println("エラーハンドラー")

				if err := systemErrorHandlerPresenter.ErrorResponse(cctx.Ctx); err != nil {
					return xerrors.Errorf("systemErrorHandlerPresenter.ErrorResponse(): %w", err)
				}

				return err
			}

			return nil
		}
	}
}
