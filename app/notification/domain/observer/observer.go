package observer

import "github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"

type Subscriber[T event.DomainEvent] interface {
	Subscribe(event T)
}

type Publisher[T event.DomainEvent] interface {
	Register(subscribers ...Subscriber[T])
	Remove(subscriber Subscriber[T])
	Publish(event T)
}
