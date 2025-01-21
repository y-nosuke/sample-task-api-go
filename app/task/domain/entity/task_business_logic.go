package entity

import (
	"time"

	"golang.org/x/xerrors"

	"github.com/google/uuid"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

// TODO タスク完了イベント・タスク未完了イベントの発行・登録

func (t *Task) Edit(title string, detail *string, completed bool, deadline *time.Time, userID uuid.UUID) (*event.TaskEdited, error) {
	t.title = title
	t.detail = detail
	t.completed = completed
	t.deadline = deadline
	t.editedBy = userID
	t.editedAt = time.Now()

	updated, err := taskEdited(t)
	if err != nil {
		return nil, xerrors.Errorf("taskEdited(): %w", err)
	}

	return updated, nil
}

func (t *Task) Complete(userID uuid.UUID) (*event.TaskCompleted, error) {
	t.completed = true
	t.editedBy = userID
	t.editedAt = time.Now()

	completed, err := taskCompleted(t)
	if err != nil {
		return nil, xerrors.Errorf("taskCompleted(): %w", err)
	}

	return completed, nil
}

func (t *Task) UnComplete(userID uuid.UUID) (*event.TaskUnCompleted, error) {
	t.completed = false
	t.editedBy = userID
	t.editedAt = time.Now()

	unCompleted, err := taskUnCompleted(t)
	if err != nil {
		return nil, xerrors.Errorf("taskUnCompleted(): %w", err)
	}

	return unCompleted, nil
}

// CreateEvent

func taskEdited(task *Task) (*event.TaskEdited, error) {
	updated, err := event.NewTaskEdited(task.id, task.title, task.detail, task.completed, task.deadline, task.editedBy, task.editedAt)
	if err != nil {
		return nil, xerrors.Errorf("event.NewTaskEdited(): %w", err)
	}
	return updated, nil
}

func taskCompleted(task *Task) (*event.TaskCompleted, error) {
	completed, err := event.NewTaskCompleted(task.id, task.editedBy, task.editedAt)
	if err != nil {
		return nil, xerrors.Errorf("event.NewTaskCompleted(): %w", err)
	}
	return completed, nil
}

func taskUnCompleted(task *Task) (*event.TaskUnCompleted, error) {
	unCompleted, err := event.NewTaskUnCompleted(task.id, task.editedBy, task.editedAt)
	if err != nil {
		return nil, xerrors.Errorf("event.NewTaskUnCompleted(): %w", err)
	}
	return unCompleted, nil
}
