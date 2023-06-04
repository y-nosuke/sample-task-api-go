package infrastructure

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	fauth "github.com/y-nosuke/sample-task-api-go/framework/auth/interfaces"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	fdatabase "github.com/y-nosuke/sample-task-api-go/framework/database/interfaces"
	ferrors "github.com/y-nosuke/sample-task-api-go/framework/errors/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
	"github.com/y-nosuke/sample-task-api-go/task/application/usecases"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/controllers"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/database"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/presenters"
	"io"
	"strings"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Validator = controllers.NewValidator()

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

	taskRepository := database.NewTaskRepository()
	taskPresenter := presenters.NewTaskPresenter()
	registerTaskUseCase := usecases.NewRegisterTaskUseCase(taskRepository, taskPresenter)
	getAllTaskUseCase := usecases.NewGetAllTaskUseCase(taskRepository, taskPresenter)
	getTaskUseCase := usecases.NewGetTaskUseCase(taskRepository, taskPresenter)
	updateTaskUseCase := usecases.NewUpdateTaskUseCase(taskRepository, taskPresenter)
	completeTaskUseCase := usecases.NewCompleteTaskUseCase(taskRepository, taskPresenter)
	unCompleteTaskUseCase := usecases.NewUnCompleteTaskUseCase(taskRepository, taskPresenter)
	deleteTaskUseCase := usecases.NewDeleteTaskUseCase(taskRepository, taskPresenter)
	taskController := controllers.NewTaskController(
		registerTaskUseCase,
		getAllTaskUseCase,
		getTaskUseCase,
		updateTaskUseCase,
		completeTaskUseCase,
		unCompleteTaskUseCase,
		deleteTaskUseCase,
	)

	openapi.RegisterHandlers(g, taskController)

	// ここで処理しないとjaegerのtracingが取れなくなる
	e.Logger.Fatal(e.Start(":1323"))
	return e
}

func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics")
}
