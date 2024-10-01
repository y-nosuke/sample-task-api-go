package factory

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"time"
)

func CreateTask(title string, detail *string, deadline *time.Time) *entity.Task {
	id := uuid.New()
	return entity.NewTask(id, title, detail, false, deadline, uuid.Nil, time.Time{}, uuid.Nil, time.Time{}, uuid.Nil)
}
