package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskCreated struct {
	TaskEvent
	Data TaskCreatedData
}

type TaskCreatedData struct {
	Title     string     `json:"title"`
	Detail    *string    `json:"detail"`
	Completed bool       `json:"completed"`
	Deadline  *time.Time `json:"deadline"`
	CreatedBy *uuid.UUID `json:"created_by"`
	CreatedAt *time.Time `json:"created_at"`
}

func NewTaskCreated(task *entity.Task) *TaskCreated {
	return &TaskCreated{
		TaskEvent: newTaskEvent(task.Id, ETaskCreated),
		Data: TaskCreatedData{
			Title:     task.Title,
			Detail:    task.Detail,
			Completed: task.Completed,
			Deadline:  task.Deadline,
			CreatedBy: task.CreatedBy,
			CreatedAt: task.CreatedAt,
		},
	}
}
