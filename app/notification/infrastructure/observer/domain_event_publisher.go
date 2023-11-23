package observer

import (
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/observer"
)

type DomainEventPublisherImpl[T event.DomainEvent] struct {
	subscribers []observer.Subscriber[T]
}

func NewDomainEventPublisherImpl() *DomainEventPublisherImpl[event.DomainEvent] {
	return &DomainEventPublisherImpl[event.DomainEvent]{}
}

func (p *DomainEventPublisherImpl[T]) Register(subscribers ...observer.Subscriber[T]) {
	for _, s := range subscribers {
		p.subscribers = append(p.subscribers, s)
	}
}

func (p *DomainEventPublisherImpl[T]) Remove(subscriber observer.Subscriber[T]) {
	for i, s := range p.subscribers {
		if s == subscriber {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
		}
	}
}

func (p *DomainEventPublisherImpl[T]) Publish(event T) {
	for _, s := range p.subscribers {
		s.Subscribe(event)
	}
}
