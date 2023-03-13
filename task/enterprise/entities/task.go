package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	Title     string
	Detail    string
	Completed bool
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   uuid.UUID
}

func NewTask(title string, detail string, deadline time.Time) (*Task, error) {
	id := uuid.New()

	return &Task{Id: id, Title: title, Detail: detail, Completed: false, Deadline: deadline}, nil
}
