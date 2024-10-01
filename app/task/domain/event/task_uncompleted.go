package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
)

type TaskUnCompleted struct {
	TaskEventCommon
	data TaskUnCompletedData
}

type TaskUnCompletedData struct {
	UpdatedBy uuid.UUID `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskUnCompleted(task *entity.Task) *TaskUnCompleted {
	return &TaskUnCompleted{
		TaskEventCommon: *newTaskEventCommon(task.Id()),
		data: TaskUnCompletedData{
			UpdatedBy: task.CreatedBy,
			UpdatedAt: task.CreatedAt,
		},
	}
}

func (t TaskUnCompleted) Data() any {
	return t.data
}
