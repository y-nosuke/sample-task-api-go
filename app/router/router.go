package router

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure"
	fap "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure/presenter"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	fdatabase "github.com/y-nosuke/sample-task-api-go/app/framework/database/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors/infrastructure"
	fep "github.com/y-nosuke/sample-task-api-go/app/framework/errors/infrastructure/presenter"
	tr "github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/router"
	"io"
	"strings"
)

func Router() *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Validator = NewValidator()

	c := jaegertracing.New(e, urlSkipper)
	defer func(c io.Closer) {
		// TODO: deferのエラー処理
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		echoprometheus.NewMiddleware("sample_task_api_go"),
	)

	e.GET("/metrics", echoprometheus.NewHandler())

	g := e.Group("/api/v1")

	systemErrorHandlerPresenterImpl := fep.NewSystemErrorHandlerPresenterImpl()
	authHandlerPresenterImpl := fap.NewAuthHandlerPresenterImpl()
	g.Use(
		fcontext.CustomContextMiddleware,
		ferrors.ErrorHandlerMiddleware(systemErrorHandlerPresenterImpl),
		fauth.ValidateTokenMiddleware(authHandlerPresenterImpl),
		fdatabase.TransactionMiddleware,
	)

	tr.TaskRouter(g)

	// ここで処理しないとjaegerのtracingが取れなくなる
	e.Logger.Fatal(e.Start(":1323"))
	return e
}

func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics")
}
