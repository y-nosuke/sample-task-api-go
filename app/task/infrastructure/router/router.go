package router

import (
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/handler"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/repository"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

func TaskRouter(g *echo.Group) {
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
		taskPresenterImpl,
	)

	openapi.RegisterHandlers(g, taskHandler)
}
