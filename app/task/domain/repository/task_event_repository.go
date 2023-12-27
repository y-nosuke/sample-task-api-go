package repository

import (
	"context"

	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

type TaskEventRepository interface {
	Register(context.Context, event.TaskEvent[event.TaskCreatedData]) error
}
