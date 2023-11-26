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
	taskEventRepositoryImpl := repository.NewTaskEventRepositoryImpl()
	taskPresenterImpl := presenter.NewTaskPresenterImpl()
	registerTaskUseCase := usecase.NewRegisterTaskUseCase(taskRepositoryImpl, taskEventRepositoryImpl, taskPresenterImpl, publisher)
	getAllTaskUseCase := usecase.NewGetAllTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	getTaskUseCase := usecase.NewGetTaskUseCase(taskRepositoryImpl, taskPresenterImpl)
	updateTaskUseCase := usecase.NewUpdateTaskUseCase(taskRepositoryImpl, taskEventRepositoryImpl, taskPresenterImpl, publisher)
	completeTaskUseCase := usecase.NewCompleteTaskUseCase(taskRepositoryImpl, taskEventRepositoryImpl, taskPresenterImpl, publisher)
	unCompleteTaskUseCase := usecase.NewUnCompleteTaskUseCase(taskRepositoryImpl, taskEventRepositoryImpl, taskPresenterImpl, publisher)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCase(taskRepositoryImpl, taskEventRepositoryImpl, taskPresenterImpl, publisher)
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
