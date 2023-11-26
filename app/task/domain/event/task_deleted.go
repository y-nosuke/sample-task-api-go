package event

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

type TaskDeleted struct {
	TaskEventCommon
	data TaskDeletedData
}

type TaskDeletedData struct {
	DeletedBy *uuid.UUID `json:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewTaskDeleted(task *entity.Task, deletedBy *uuid.UUID) *TaskDeleted {
	now := time.Now()
	return &TaskDeleted{
		TaskEventCommon: *newTaskEventCommon(task.Id, ETaskDeleted),
		data: TaskDeletedData{
			DeletedBy: deletedBy,
			DeletedAt: &now,
		},
	}
}

func (t TaskDeleted) Data() any {
	return t.data
}
