package event

import (
	"github.com/google/uuid"
)

type DomainEvent interface {
	ID() uuid.UUID
}

type DomainEventImpl struct {
	id uuid.UUID
}

func NewDomainEvent() DomainEvent {
	return &DomainEventImpl{uuid.New()}
}

func (e DomainEventImpl) ID() uuid.UUID {
	return e.id
}
