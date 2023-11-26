package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskCompleted struct {
	TaskEvent
	Data TaskCompletedData
}

type TaskCompletedData struct {
	UpdatedBy *uuid.UUID `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewTaskCompleted(task *entity.Task) *TaskCompleted {
	return &TaskCompleted{
		TaskEvent: newTaskEvent(task.Id, ETaskCompleted),
		Data: TaskCompletedData{
			UpdatedBy: task.UpdatedBy,
			UpdatedAt: task.UpdatedAt,
		},
	}
}
