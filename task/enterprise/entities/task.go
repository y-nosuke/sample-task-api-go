package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
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

func NewTask(title string, detail string, deadline string) (*Task, error) {
	id := uuid.New()
	var layout = "2006-01-02"
	var parsedDeadline time.Time
	if deadline != "" {
		var err error
		parsedDeadline, err = time.Parse(layout, deadline)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
	}

	return &Task{Id: id, Title: title, Detail: detail, Completed: false, Deadline: parsedDeadline}, nil
}
