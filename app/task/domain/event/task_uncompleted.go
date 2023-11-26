package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskUnCompleted struct {
	TaskEvent
	Data TaskUnCompletedData
}

type TaskUnCompletedData struct {
	UpdatedBy *uuid.UUID `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewTaskUnCompleted(task *entity.Task) *TaskUnCompleted {
	return &TaskUnCompleted{
		TaskEvent: newTaskEvent(task.Id, ETaskUnCompleted),
		Data: TaskUnCompletedData{
			UpdatedBy: task.CreatedBy,
			UpdatedAt: task.CreatedAt,
		},
	}
}
