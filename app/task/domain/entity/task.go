package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaskSlice []*Task

type Task struct {
	id        uuid.UUID
	title     string
	detail    *string
	completed bool
	deadline  *time.Time
	createdBy uuid.UUID
	createdAt time.Time
	editedBy  uuid.UUID
	editedAt  time.Time
	Version   uuid.UUID
}

func (t *Task) Id() uuid.UUID {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Detail() *string {
	return t.detail
}

func (t *Task) Completed() bool {
	return t.completed
}

func (t *Task) Deadline() *time.Time {
	return t.deadline
}

func (t *Task) CreatedBy() uuid.UUID {
	return t.createdBy
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) EditedBy() uuid.UUID {
	return t.editedBy
}

func (t *Task) EditedAt() time.Time {
	return t.editedAt
}

func NewTask(id uuid.UUID, title string, detail *string, completed bool, deadline *time.Time, createdBy uuid.UUID, createdAt time.Time, editedBy uuid.UUID, editedAt time.Time, version uuid.UUID) *Task {
	return &Task{id: id, title: title, detail: detail, completed: completed, deadline: deadline, createdBy: createdBy, createdAt: createdAt, editedBy: editedBy, editedAt: editedAt, Version: version}
}
