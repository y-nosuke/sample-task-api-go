package event

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type DomainEvent interface {
	ID() uuid.UUID
}

type DomainEventImpl struct {
	id uuid.UUID
}

func NewDomainEvent() (DomainEvent, error) {
	u, err := uuid.NewV7()
	if err != nil {
		return nil, xerrors.Errorf("uuid.NewV7(): %w", err)
	}
	return &DomainEventImpl{u}, nil
}

func (e DomainEventImpl) ID() uuid.UUID {
	return e.id
}
