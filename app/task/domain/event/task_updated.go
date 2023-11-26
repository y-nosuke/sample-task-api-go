package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskUpdated struct {
	TaskEvent
	Data TaskUpdatedData
}

type TaskUpdatedData struct {
	Title     string     `json:"title"`
	Detail    *string    `json:"detail"`
	Completed bool       `json:"completed"`
	Deadline  *time.Time `json:"deadline"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewTaskUpdated(task *entity.Task) *TaskUpdated {
	return &TaskUpdated{
		TaskEvent: newTaskEvent(task.Id, ETaskUpdated),
		Data: TaskUpdatedData{
			Title:     task.Title,
			Detail:    task.Detail,
			Completed: task.Completed,
			Deadline:  task.Deadline,
			UpdatedBy: task.UpdatedBy,
			UpdatedAt: task.UpdatedAt,
		},
	}
}
