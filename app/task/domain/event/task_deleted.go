package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskDeleted struct {
	TaskEvent
	Data TaskDeletedData
}

type TaskDeletedData struct {
	DeletedBy *uuid.UUID `json:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewTaskDeleted(task *entity.Task, deletedBy *uuid.UUID) *TaskDeleted {
	now := time.Now()
	return &TaskDeleted{
		TaskEvent: newTaskEvent(task.Id, ETaskDeleted),
		Data: TaskDeletedData{
			DeletedBy: deletedBy,
			DeletedAt: &now,
		},
	}
}
