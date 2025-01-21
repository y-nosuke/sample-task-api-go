package router

import (
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/admin/infrastructure/handler"
)

func AdminRouter(g *echo.Group) {
	adminHandler := handler.NewAdminHandler()
	g.GET("/health", adminHandler.Ping)
}
