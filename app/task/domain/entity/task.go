package entity

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
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedBy uuid.UUID
	UpdatedAt time.Time
	Version   uuid.UUID
}

type TaskSlice []*Task

func NewTask(title string, detail *string, deadline *time.Time) *Task {
	id := uuid.New()
	return &Task{Id: id, Title: title, Detail: detail, Completed: false, Deadline: deadline}
}

func (t *Task) Update(title string, detail *string, deadline *time.Time, version uuid.UUID) {
	t.Title = title
	t.Detail = detail
	t.Deadline = deadline
	t.Version = version
}

func (t *Task) Complete(version uuid.UUID) {
	t.Completed = true
	t.Version = version
}

func (t *Task) UnComplete(version uuid.UUID) {
	t.Completed = false
	t.Version = version
}
