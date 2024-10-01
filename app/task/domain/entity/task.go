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
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedBy uuid.UUID
	UpdatedAt time.Time
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

func NewTask(id uuid.UUID, title string, detail *string, completed bool, deadline *time.Time, createdBy uuid.UUID, createdAt time.Time, updatedBy uuid.UUID, updatedAt time.Time, version uuid.UUID) *Task {
	return &Task{id: id, title: title, detail: detail, completed: completed, deadline: deadline, CreatedBy: createdBy, CreatedAt: createdAt, UpdatedBy: updatedBy, UpdatedAt: updatedAt, Version: version}
}
