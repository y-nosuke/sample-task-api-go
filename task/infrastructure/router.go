package infrastructure

import (
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	fdatabase "github.com/y-nosuke/sample-task-api-go/framework/database/interfaces"
	ferror "github.com/y-nosuke/sample-task-api-go/framework/error/interfaces"
	"github.com/y-nosuke/sample-task-api-go/generated/interfaces/openapi"
	"github.com/y-nosuke/sample-task-api-go/task/application/usecases"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/controllers"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/database"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/presenters"
)

func Router() *echo.Echo {
	e := echo.New()

	swagger, err := openapi.GetSwagger("/api/v1")
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		fcontext.CustomContextMiddleware,
		ferror.ErrorHandlerMiddleware,
		fdatabase.TransactionMiddleware,
	)

	g := e.Group("/api/v1")

	g.Use(oapiMiddleware.OapiRequestValidator(swagger))

	taskRepository := database.NewTaskRepository()
	taskPresenter := presenters.NewTaskPresenter()
	registerTaskUseCase := usecases.NewRegisterTaskUseCase(taskRepository, taskPresenter)
	getAllTaskUseCase := usecases.NewGetAllTaskUseCase(taskRepository, taskPresenter)
	getTaskUseCase := usecases.NewGetTaskUseCase(taskRepository, taskPresenter)
	updateTaskUseCase := usecases.NewUpdateTaskUseCase(taskRepository, taskPresenter)
	deleteTaskUseCase := usecases.NewDeleteTaskUseCase(taskRepository, taskPresenter)
	taskController := controllers.NewTaskController(registerTaskUseCase, getAllTaskUseCase, getTaskUseCase, updateTaskUseCase, deleteTaskUseCase)

	openapi.RegisterHandlers(g, taskController)

	return e
}
