package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"golang.org/x/xerrors"
)

func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		cctx := fcontext.NewCustomContext(ectx)
		fmt.Println("ContextMiddleware start. コンテキストを設定します。")

		if err := next(cctx); err != nil {
			return xerrors.Errorf("next(): %w", err)
		}

		return nil
	}
}
