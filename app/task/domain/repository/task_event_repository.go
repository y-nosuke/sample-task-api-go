package repository

import (
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

type TaskEventRepository interface {
	Register(fcontext.Context, event.TaskEvent) error
}
