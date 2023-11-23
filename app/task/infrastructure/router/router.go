package router

import (
	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
	"github.com/y-nosuke/sample-task-api-go/app/task/application/usecase"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/handler"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/presenter"
	"github.com/y-nosuke/sample-task-api-go/app/task/infrastructure/repository"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

func TaskRouter(g *echo.Group, publisher observer.Publisher[event.DomainEvent]) {
	taskRepositoryImpl := repository.NewTaskRepositoryImpl()
	taskPresenterImpl := presenter.NewTaskPresenterImpl()
	registerTaskUseCase := usecase.NewRegisterTaskUseCase(taskRepositoryImpl, taskPresenterImpl, publisher)
	getAllTaskUseCase := usecase.NewGetAllTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	getTaskUseCase := usecase.NewGetTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	updateTaskUseCase := usecase.NewUpdateTaskUseCase(taskRepositoryImpl, taskPresenterImpl, publisher)
	completeTaskUseCase := usecase.NewCompleteTaskUseCase(taskRepositoryImpl, taskPresenterImpl, publisher)
	unCompleteTaskUseCase := usecase.NewUnCompleteTaskUseCase(taskRepositoryImpl, taskPresenterImpl, publisher)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCase(taskRepositoryImpl, taskPresenterImpl, publisher)
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
