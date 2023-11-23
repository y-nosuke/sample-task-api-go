package router

import (
	"fmt"
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
	"github.com/y-nosuke/sample-task-api-go/app/notification/infrastructure/observer"
	tr "github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/router"
	"io"
	"os"
	"strconv"
	"strings"
)

func Router() (e *echo.Echo, err error) {
	e = echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Validator = NewValidator()

	c := jaegertracing.New(e, urlSkipper)
	defer func(c io.Closer) {
		if closeErr := c.Close(); closeErr != nil {
			err = fmt.Errorf("original error: %v, defer close error: %v", err, closeErr)
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

	domainEventPublisherImpl := observer.NewDomainEventPublisherImpl()
	slackSubscriberImpl := observer.NewSlackSubscriberImpl(os.Getenv("SLACK_TOKEN"), os.Getenv("CHANNEL_ID"))
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi(): %v", err)
	}
	mailSubscriberImpl := observer.NewMailSubscriberImpl(os.Getenv("MAIL_HOST"), port, os.Getenv("MAIL_FROM"), os.Getenv("MAIL_TO"))
	domainEventPublisherImpl.Register(slackSubscriberImpl, mailSubscriberImpl)

	tr.TaskRouter(g, domainEventPublisherImpl)

	// ここで処理しないとjaegerのtracingが取れなくなる
	e.Logger.Fatal(e.Start(":1323"))
	return e, nil
}

func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics")
}
