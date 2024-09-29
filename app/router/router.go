package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	fep "github.com/y-nosuke/sample-task-api-go/app/framework/io/infrastructure/presenter"
	fmiddleware "github.com/y-nosuke/sample-task-api-go/app/framework/middleware"
	"github.com/y-nosuke/sample-task-api-go/app/framework/validation"
	"github.com/y-nosuke/sample-task-api-go/app/notification/infrastructure/observer"
	tr "github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/router"
	"go.opentelemetry.io/otel"

	//nolint:staticcheck
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"os"
	"strconv"
	"strings"
	"time"
)

// golangci-lintでエラーになるので一時的にコメントアウト
//var tracer trace.Tracer

func Router() (e *echo.Echo, err error) {
	e = echo.New()

	e.HTTPErrorHandler = ferrors.CustomHTTPErrorHandler
	e.Validator = validation.NewValidator()

	//tracer = otel.Tracer("github.com/y-nosuke/sample-task-api-go")
	otel.Tracer("github.com/y-nosuke/sample-task-api-go")

	ctx := context.Background()
	endpoint := os.Getenv("EXPORTER_ENDPOINT")
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(endpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("otlptracehttp.New(): %v", err)
	}

	r, err := resource.New(
		ctx,
		resource.WithProcessPID(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName("sample-task-api-go"),
			semconv.ServiceVersion("1.0.0"),
			semconv.DeploymentEnvironment("local"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("resource.New(): %v", err)
	}

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		otelecho.Middleware("sample-task-api-go",
			otelecho.WithSkipper(urlSkipper),
			otelecho.WithTracerProvider(sdkTrace.NewTracerProvider(
				sdkTrace.WithBatcher(exporter,
					// Default is 5s. Set to 1s for demonstrative purposes.
					sdkTrace.WithBatchTimeout(time.Second),
				),
				sdkTrace.WithResource(r),
			)),
			otelecho.WithPropagators(
				propagation.NewCompositeTextMapPropagator(
					propagation.TraceContext{},
					propagation.Baggage{},
				),
			),
		),
		echoprometheus.NewMiddleware("sample_task_api_go"),
	)

	e.GET("/metrics", echoprometheus.NewHandler())

	g := e.Group("/api/v1")

	systemErrorHandlerPresenterImpl := fep.NewSystemErrorHandlerPresenterImpl()
	authHandlerPresenterImpl := fep.NewAuthHandlerPresenterImpl()
	g.Use(
		fmiddleware.CustomContextMiddleware,
		fmiddleware.ErrorHandlerMiddleware(systemErrorHandlerPresenterImpl),
		fmiddleware.ValidateTokenMiddleware(authHandlerPresenterImpl),
		fmiddleware.TransactionMiddleware,
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
