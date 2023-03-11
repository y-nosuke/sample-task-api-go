package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	fdatabase "github.com/y-nosuke/sample-task-api-go/framework/database/interfaces"
	ferror "github.com/y-nosuke/sample-task-api-go/framework/error/interfaces"
	"github.com/y-nosuke/sample-task-api-go/task/application/usecases"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/controllers"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/database"
	"github.com/y-nosuke/sample-task-api-go/task/interfaces/presenters"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		fcontext.CustomContextMiddleware,
		ferror.ErrorHandlerMiddleware,
		fdatabase.TransactionMiddleware,
	)

	g := e.Group("/api/v1/tasks")

	taskRepository := database.NewTaskRepository()
	registerTaskUseCase := usecases.NewRegisterTaskUseCase(taskRepository)
	taskPresenter := presenters.NewTaskPresenter()
	taskController := controllers.NewTaskController(registerTaskUseCase, taskPresenter)

	g.POST("", taskController.RegisterTask)

	return e
}
