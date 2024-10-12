package entity

import (
	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"time"
)

func (t *Task) Update(title string, detail *string, deadline *time.Time, userID uuid.UUID) *event.TaskEdited {
	t.title = title
	t.detail = detail
	t.deadline = deadline
	t.editedBy = userID
	t.editedAt = time.Now()

	return taskUpdated(t)
}

func (t *Task) Complete(userID uuid.UUID) *event.TaskCompleted {
	t.completed = true
	t.editedBy = userID
	t.editedAt = time.Now()

	return taskCompleted(t)
}

func (t *Task) UnComplete(userID uuid.UUID) *event.TaskUnCompleted {
	t.completed = false
	t.editedBy = userID
	t.editedAt = time.Now()

	return taskUnCompleted(t)
}

// CreateEvent

func taskUpdated(task *Task) *event.TaskEdited {
	return event.NewTaskUpdated(task.id, task.title, task.detail, task.completed, task.deadline, task.editedBy, task.editedAt)
}

func taskCompleted(task *Task) *event.TaskCompleted {
	return event.NewTaskCompleted(task.id, task.editedBy, task.editedAt)
}

func taskUnCompleted(task *Task) *event.TaskUnCompleted {
	return event.NewTaskUnCompleted(task.id, task.editedBy, task.editedAt)
}
