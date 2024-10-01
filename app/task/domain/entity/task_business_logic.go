package entity

import (
	"github.com/google/uuid"
	"time"
)

func (t *Task) Update(title string, detail *string, deadline *time.Time, version uuid.UUID) {
	t.title = title
	t.detail = detail
	t.deadline = deadline
	t.Version = version
}

func (t *Task) Complete(version uuid.UUID) {
	t.completed = true
	t.Version = version
}

func (t *Task) UnComplete(version uuid.UUID) {
	t.completed = false
	t.Version = version
}
