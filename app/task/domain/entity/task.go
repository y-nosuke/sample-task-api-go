package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        *uuid.UUID
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
}

type TaskSlice []*Task

func NewTask(title string, detail *string, deadline *time.Time) *Task {
	id := uuid.New()
	return &Task{Id: &id, Title: title, Detail: detail, Completed: false, Deadline: deadline}
}

func (t *Task) Update(title string, detail *string, deadline *time.Time) {
	t.Title = title
	t.Detail = detail
	t.Deadline = deadline
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) UnComplete() {
	t.Completed = false
}
