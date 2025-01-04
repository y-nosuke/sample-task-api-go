package middleware

import (
	"errors"
	"fmt"
	"net/http"

	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/io/application/presenter"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"golang.org/x/xerrors"
)

func ErrorResponseHandleMiddleware(errorHandlerPresenter presenter.ErrorPresenter) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) (err error) {
			cctx := context.CastContext(ectx)
			fmt.Println("ErrorResponseHandleMiddleware start. エラーレスポンスハンドラーを実行します。")

			defer func() {
				if p := recover(); p != nil {
					if resErr := errorHandlerPresenter.InternalServerError(cctx); resErr != nil {
						panic(fmt.Sprintf("original panic: %v, errorHandlerPresenter.InternalServerError(): %v", p, resErr))
					}
					panic(fmt.Sprintf("panic occurred. %v", p))
				} else if err != nil {
					var he *echo.HTTPError
					var be *ferrors.BusinessError
					if errors.As(err, &he) {
						errorResponse := &openapi.ErrorResponse{Message: http.StatusText(he.Code)}
						if resErr := ectx.JSON(he.Code, errorResponse); resErr != nil {
							err = xerrors.Errorf("ectx.JSON(): %w", resErr)
						}
					} else if errors.As(err, &be) {
						switch be.ErrorCode() {
						case ferrors.BadRequest:
							if resErr := errorHandlerPresenter.BadRequest(cctx, "バリデーションエラーです。", err); resErr != nil {
								err = xerrors.Errorf("original error: %v, errorHandlerPresenter.BadRequest(): %w", err, resErr)
							}
						case ferrors.Unauthorized:
							if resErr := errorHandlerPresenter.Unauthorized(cctx, "認証されていません。"); resErr != nil {
								err = xerrors.Errorf("original error: %v, errorHandlerPresenter.Unauthorized(): %w", err, resErr)
							}
						case ferrors.Forbidden:
							if resErr := errorHandlerPresenter.Forbidden(cctx, "指定された操作は許可されていません。"); resErr != nil {
								err = xerrors.Errorf("original error: %v, errorHandlerPresenter.Forbidden(): %w", err, resErr)
							}
						case ferrors.NotFound:
							if resErr := errorHandlerPresenter.NotFound(cctx, "指定されたタスクが見つかりませんでした。"); resErr != nil {
								err = xerrors.Errorf("original error: %v, errorHandlerPresenter.NotFound(): %w", err, resErr)
							}
						case ferrors.Conflict:
							if resErr := errorHandlerPresenter.Conflict(cctx, "タスクは既に更新済みです。"); resErr != nil {
								err = xerrors.Errorf("original error: %v, errorHandlerPresenter.Conflict(): %w", err, resErr)
							}
						default:
							err = xerrors.Errorf("unhandled default case: %v", be.ErrorCode())
						}
					} else {
						if resErr := errorHandlerPresenter.InternalServerError(cctx); resErr != nil {
							err = xerrors.Errorf("original error: %v, errorHandlerPresenter.InternalServerError(): %w", err, resErr)
						}
					}
				}
			}()

			err = next(cctx)

			return
		}
	}
}
