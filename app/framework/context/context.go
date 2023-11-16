package context

import (
	"context"
	"github.com/labstack/echo/v4"
)

type ctxKey int

const (
	ECTX ctxKey = iota
)

func SetEctx(cctx *CustomContextImpl, ectx echo.Context) {
	cctx.WithValue(ECTX, ectx)
}

func GetEctx(ctx context.Context) echo.Context {
	return ctx.Value(ECTX).(echo.Context)
}

func Ctx(ectx echo.Context) context.Context {
	return ectx.(*CustomContextImpl).Ctx
}

func Cctx(ectx echo.Context) *CustomContextImpl {
	return ectx.(*CustomContextImpl)
}

type CustomContext interface {
	WithValue(key any, val any)
	Value(key any) any
}

type CustomContextImpl struct {
	echo.Context
	Ctx context.Context
}

func NewCustomContext(ectx echo.Context, ctx context.Context) *CustomContextImpl {
	return &CustomContextImpl{ectx, ctx}
}

func (c *CustomContextImpl) WithValue(key any, val any) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

func (c *CustomContextImpl) Value(key any) any {
	return c.Ctx.Value(key)
}
