package infrastructure

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
)

type ctxKey int

const (
	ECTX ctxKey = iota
)

func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("コンテキストを初期化します。")
		cctx := fcontext.NewCustomContext(ectx, context.Background())
		cctx.WithValue(ECTX, ectx)
		return next(cctx)
	}
}

func Cctx(ectx echo.Context) *fcontext.CustomContext {
	return ectx.(*fcontext.CustomContext)
}

func Ectx(ctx context.Context) echo.Context {
	return ctx.Value(ECTX).(echo.Context)
}
