package entity

import (
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        *uuid.UUID
	Title     string
	Detail    *string
	Completed bool
	Deadline  *time.Time
	CreatedBy *uuid.UUID
	CreatedAt *time.Time
	UpdatedBy *uuid.UUID
	UpdatedAt *time.Time
	Version   *uuid.UUID
}

type TaskSlice []*Task

func NewTask(title string, detail *string, deadline *time.Time) (*Task, *event.TaskCreated) {
	id := uuid.New()

	return &Task{Id: &id, Title: title, Detail: detail, Completed: false, Deadline: deadline}, event.NewTaskCreated(&id)
}

func (t *Task) Update(title string, detail *string, deadline *time.Time, version *uuid.UUID) *event.TaskUpdated {
	t.Title = title
	t.Detail = detail
	t.Deadline = deadline
	t.Version = version

	return event.NewTaskUpdated(t.Id)
}

func (t *Task) Complete(version *uuid.UUID) *event.TaskCompleted {
	t.Completed = true
	t.Version = version

	return event.NewTaskCompleted(t.Id)
}

func (t *Task) UnComplete(version *uuid.UUID) *event.TaskUnCompleted {
	t.Completed = false
	t.Version = version

	return event.NewTaskUnCompleted(t.Id)
}
