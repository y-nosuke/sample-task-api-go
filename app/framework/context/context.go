package context

import (
	"context"
	"github.com/labstack/echo/v4"
)

type ctxKey int

const (
	ECTX ctxKey = iota
)

func SetEctx(cctx *CustomContext, ectx echo.Context) {
	cctx.WithValue(ECTX, ectx)
}

func GetEctx(ctx context.Context) echo.Context {
	return ctx.Value(ECTX).(echo.Context)
}

func Ctx(ectx echo.Context) context.Context {
	return ectx.(*CustomContext).Ctx
}

func Cctx(ectx echo.Context) *CustomContext {
	return ectx.(*CustomContext)
}

type CustomContext struct {
	echo.Context
	Ctx context.Context
}

func NewCustomContext(ectx echo.Context, ctx context.Context) *CustomContext {
	return &CustomContext{ectx, ctx}
}

func (c *CustomContext) WithValue(key any, val any) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

func (c *CustomContext) Value(key any) any {
	return c.Ctx.Value(key)
}
