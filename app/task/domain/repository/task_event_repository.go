package repository

import (
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

type TaskEventRepository interface {
	RegisterTaskCreated(fcontext.Context, *event.TaskCreated) error
	RegisterTaskUpdated(fcontext.Context, *event.TaskEdited) error
	RegisterTaskCompleted(fcontext.Context, *event.TaskCompleted) error
	RegisterTaskUnCompleted(fcontext.Context, *event.TaskUnCompleted) error
	RegisterTaskDeleted(fcontext.Context, *event.TaskDeleted) error
}
