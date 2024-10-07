package middleware

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"github.com/y-nosuke/sample-task-api-go/app/framework/io/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"golang.org/x/xerrors"
	"net/http"
)

func ErrorResponseHandleMiddleware(systemErrorHandlerPresenter presenter.SystemErrorPresenter) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) (err error) {
			cctx := context.CastContext(ectx)
			fmt.Println("ErrorResponseHandleMiddleware start. エラーレスポンスハンドラーを実行します。")

			defer func() {
				if p := recover(); p != nil {
					if resErr := systemErrorHandlerPresenter.InternalServerError(cctx); resErr != nil {
						panic(fmt.Sprintf("original panic: %v, systemErrorHandlerPresenter.InternalServerError(): %v", p, resErr))
					}
					panic(fmt.Sprintf("panic occurred. %v", p))
				} else if err != nil {
					var he *echo.HTTPError
					if errors.As(err, &he) {
						errorResponse := &openapi.ErrorResponse{Message: http.StatusText(he.Code)}
						if resErr := ectx.JSON(he.Code, errorResponse); resErr != nil {
							err = xerrors.Errorf("ectx.JSON(): %w", resErr)
							return
						}
					} else if ferrors.IsSystemError(err) {
						if resErr := systemErrorHandlerPresenter.InternalServerError(cctx); resErr != nil {
							err = xerrors.Errorf("systemErrorHandlerPresenter.InternalServerError() original err=%v: %w", err, resErr)
							return
						}
					}
					err = xerrors.Errorf("next(): %w", err)
				}
			}()

			err = next(cctx)

			return
		}
	}
}
