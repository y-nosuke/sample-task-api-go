package router

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ar "github.com/y-nosuke/sample-task-api-go/app/admin/infrastructure/router"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	fep "github.com/y-nosuke/sample-task-api-go/app/framework/io/infrastructure/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/framework/mail"
	fmiddleware "github.com/y-nosuke/sample-task-api-go/app/framework/middleware"
	fotel "github.com/y-nosuke/sample-task-api-go/app/framework/otel"
	"github.com/y-nosuke/sample-task-api-go/app/framework/slack"
	"github.com/y-nosuke/sample-task-api-go/app/framework/validation"
	"github.com/y-nosuke/sample-task-api-go/app/notification/infrastructure/observer"
	tr "github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/router"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// golangci-lintでエラーになるので一時的にコメントアウト
// var tracer trace.Tracer

func Router() (e *echo.Echo, err error) {
	e = echo.New()

	e.HTTPErrorHandler = ferrors.CustomHTTPErrorHandler
	e.Validator = validation.NewValidator()

	// tracer = otel.Tracer("github.com/y-nosuke/sample-task-api-go")
	otel.Tracer("github.com/y-nosuke/sample-task-api-go")

	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(fotel.Cfg.ExporterEndpoint), otlptracehttp.WithInsecure())
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
		middleware.LoggerWithConfig(middleware.LoggerConfig{Skipper: urlSkipper}),
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

	errorPresenterImpl := fep.NewErrorPresenterImpl()
	g.Use(
		fmiddleware.ContextMiddleware,
		fmiddleware.ErrorLogHandleMiddleware,
		fmiddleware.ErrorResponseHandleMiddleware(errorPresenterImpl),
		fmiddleware.ValidateTokenMiddleware(),
		fmiddleware.TransactionMiddleware,
	)

	domainEventPublisherImpl := observer.NewDomainEventPublisherImpl()
	slackSubscriberImpl := observer.NewSlackSubscriberImpl(slack.Cfg.SlackToken, slack.Cfg.ChannelID)
	mailSubscriberImpl := observer.NewMailSubscriberImpl(mail.Cfg.Host, mail.Cfg.Port, mail.Cfg.From, mail.Cfg.To)
	domainEventPublisherImpl.Register(slackSubscriberImpl, mailSubscriberImpl)

	tr.TaskRouter(g, domainEventPublisherImpl)

	admin := e.Group("/admin")

	ar.AdminRouter(admin)

	// ここで処理しないとjaegerのtracingが取れなくなる
	e.Logger.Fatal(e.Start(":1323"))
	return e, nil
}

func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics")
}
