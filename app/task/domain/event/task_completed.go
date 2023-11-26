package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskCompleted struct {
	TaskEventCommon
	data TaskCompletedData
}

type TaskCompletedData struct {
	UpdatedBy *uuid.UUID `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewTaskCompleted(task *entity.Task) *TaskCompleted {
	return &TaskCompleted{
		TaskEventCommon: *newTaskEventCommon(task.Id, ETaskCompleted),
		data: TaskCompletedData{
			UpdatedBy: task.UpdatedBy,
			UpdatedAt: task.UpdatedAt,
		},
	}
}

func (t TaskCompleted) Data() any {
	return t.data
}
