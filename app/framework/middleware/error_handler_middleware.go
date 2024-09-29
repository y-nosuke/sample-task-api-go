package middleware

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/io/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"golang.org/x/xerrors"
	"net/http"
)

func ErrorHandlerMiddleware(systemErrorHandlerPresenter presenter.SystemErrorHandlerPresenter) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			fmt.Println("エラーハンドラーを実行します。")
			cctx := context.Cctx(ectx)

			if err := next(ectx); err != nil {
				fmt.Println("エラーハンドラー")
				var he *echo.HTTPError
				if errors.As(err, &he) {
					errorResponse := &openapi.ErrorResponse{Message: http.StatusText(he.Code)}
					if err = ectx.JSON(he.Code, errorResponse); err != nil {
						return xerrors.Errorf("ectx.JSON(): %w", err)
					}
				} else {
					if err = systemErrorHandlerPresenter.ErrorResponse(cctx.Ctx); err != nil {
						return xerrors.Errorf("systemErrorHandlerPresenter.ErrorResponse(): %w", err)
					}
				}

				return err
			}

			return nil
		}
	}
}
