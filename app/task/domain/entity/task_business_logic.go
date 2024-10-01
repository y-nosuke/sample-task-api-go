package entity

import (
	"time"
)

func (t *Task) Update(title string, detail *string, deadline *time.Time) {
	t.title = title
	t.detail = detail
	t.deadline = deadline
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) UnComplete() {
	t.completed = false
}
