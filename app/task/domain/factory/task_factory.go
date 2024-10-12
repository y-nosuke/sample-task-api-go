package factory

import (
	"time"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

func CreateTask(title string, detail *string, deadline *time.Time, userID uuid.UUID) (*entity.Task, *event.TaskCreated) {
	id := uuid.New()
	now := time.Now()
	task := entity.NewTask(id, title, detail, false, deadline, userID, now, userID, now, uuid.Nil)
	return task, taskCreated(task)
}

func taskCreated(task *entity.Task) *event.TaskCreated {
	return event.NewTaskCreated(task.Id(), task.Title(), task.Detail(), task.Completed(), task.Deadline(), task.CreatedBy(), task.CreatedAt())
}
