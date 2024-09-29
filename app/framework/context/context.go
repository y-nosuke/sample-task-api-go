package context

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Context interface {
	Set(key string, val any)
	Get(key string) any
	GetContext() context.Context
}

type CustomContext struct {
	echo.Context
}

func NewCustomContext(ectx echo.Context) *CustomContext {
	return &CustomContext{ectx}
}

func (c *CustomContext) Set(key string, val any) {
	c.Context.Set(key, val)
}

func (c *CustomContext) Get(key string) any {
	return c.Context.Get(key)
}

func (c *CustomContext) GetContext() context.Context {
	return c.Context.Request().Context()
}

func GetEchoContext(ctx Context) echo.Context {
	return ctx.(*CustomContext).Context
}

func CastContext(ectx echo.Context) *CustomContext {
	return ectx.(*CustomContext)
}
