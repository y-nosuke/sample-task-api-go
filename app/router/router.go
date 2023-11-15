package router

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth/infrastructure"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	fdatabase "github.com/y-nosuke/sample-task-api-go/app/framework/database/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors/infrastructure"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/handler"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/repository"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
	"io"
	"strings"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Validator = handler.NewValidator()

	c := jaegertracing.New(e, urlSkipper)
	defer func(c io.Closer) {
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

	g.Use(
		fcontext.CustomContextMiddleware,
		ferrors.ErrorHandlerMiddleware,
		fauth.ValidateTokenMiddleware,
		fdatabase.TransactionMiddleware,
	)

	taskRepositoryImpl := repository.NewTaskRepositoryImpl()
	taskPresenterImpl := presenter.NewTaskPresenterImpl()
	registerTaskUseCase := usecase.NewRegisterTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	getAllTaskUseCase := usecase.NewGetAllTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	getTaskUseCase := usecase.NewGetTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	updateTaskUseCase := usecase.NewUpdateTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	completeTaskUseCase := usecase.NewCompleteTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	unCompleteTaskUseCase := usecase.NewUnCompleteTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	taskHandler := handler.NewTaskHandler(
		registerTaskUseCase,
		getAllTaskUseCase,
		getTaskUseCase,
		updateTaskUseCase,
		completeTaskUseCase,
		unCompleteTaskUseCase,
		deleteTaskUseCase,
	)

	openapi.RegisterHandlers(g, taskHandler)

	// ここで処理しないとjaegerのtracingが取れなくなる
	e.Logger.Fatal(e.Start(":1323"))
	return e
}

func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics")
}
