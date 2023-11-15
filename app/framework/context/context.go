package context

import (
	"context"
	"github.com/labstack/echo/v4"
)

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
