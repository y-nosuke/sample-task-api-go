package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
)

func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("コンテキストを初期化します。")
		cctx := fcontext.NewCustomContext(ectx)
		return next(cctx)
	}
}
