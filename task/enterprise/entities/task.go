package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Version   *uuid.UUID
}

func NewTask(title string, detail *string, deadline *time.Time) *Task {
	id := uuid.New()

	return &Task{Id: id, Title: title, Detail: detail, Completed: false, Deadline: deadline}
}

func (t *Task) Update(title string, detail *string, completed bool, deadline *time.Time, version *uuid.UUID) {
	t.Title = title
	t.Detail = detail
	t.Completed = completed
	t.Deadline = deadline
	t.Version = version
}