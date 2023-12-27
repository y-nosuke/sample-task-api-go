package interfaces

import (
	"errors"
	"fmt"

	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors/application/presenter"
	"golang.org/x/xerrors"
)

func ErrorHandlerMiddleware(systemErrorHandlerPresenter presenter.SystemErrorHandlerPresenter) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			fmt.Println("エラーハンドラーを実行します。")
			cctx := context.Cctx(ectx)

			if err := next(ectx); err != nil {
				fmt.Println("エラーハンドラー")

				var businessError ferrors.BusinessError
				if !errors.Is(err, &businessError) {
					fmt.Println(11)
					if err := systemErrorHandlerPresenter.ErrorResponse(cctx.Ctx); err != nil {
						return xerrors.Errorf("systemErrorHandlerPresenter.ErrorResponse(): %w", err)
					}
				}
				return err
			}

			return nil
		}
	}
}
