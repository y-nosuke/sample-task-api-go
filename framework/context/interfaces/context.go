package interfaces

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type ctxKey int

const (
	ECTX ctxKey = iota
)

type CustomContext struct {
	echo.Context
	Ctx context.Context
}

func newCustomContext(ectx echo.Context, ctx context.Context) *CustomContext {
	return &CustomContext{ectx, ctx}
}

func (c *CustomContext) WithValue(key any, val any) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

func (c *CustomContext) Value(key any) any {
	return c.Ctx.Value(key)
}

func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("コンテキストを初期化します。")
		cctx := newCustomContext(ectx, context.Background())
		cctx.WithValue(ECTX, ectx)
		return next(cctx)
	}
}

func Cctx(ectx echo.Context) *CustomContext {
	return ectx.(*CustomContext)
}

func Ectx(ctx context.Context) echo.Context {
	return ctx.Value(ECTX).(echo.Context)
}
