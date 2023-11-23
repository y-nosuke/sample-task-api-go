package observer

import (
	"fmt"
	"github.com/wneessen/go-mail"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	tevent "github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

type MailSubscriberImpl struct {
	host string
	port mail.Option
	from string
	to   string
}

func NewMailSubscriberImpl(host string, port int, from string, to string) *MailSubscriberImpl {
	_port := mail.WithPort(port)
	return &MailSubscriberImpl{host, _port, from, to}
}

func (s MailSubscriberImpl) Subscribe(event event.DomainEvent) {
	var message string
	if e, ok := event.(*tevent.TaskCreated); ok {
		message = "task created! taskID: " + e.TaskID().String()
	} else if e, ok := event.(*tevent.TaskUpdated); ok {
		message = "task updated! taskID: " + e.TaskID().String()
	} else if e, ok := event.(*tevent.TaskCompleted); ok {
		message = "task completed! taskID: " + e.TaskID().String()
	} else if e, ok := event.(*tevent.TaskUnCompleted); ok {
		message = "task uncompleted! taskID: " + e.TaskID().String()
	} else if e, ok := event.(*tevent.TaskDeleted); ok {
		message = "task deleted! taskID: " + e.TaskID().String()
	} else {
		fmt.Printf("unknown event. eventID: %s\n", event.ID())
		return
	}

	m := mail.NewMsg()
	if err := m.From(s.from); err != nil {
		fmt.Printf("failed to set From address: %s", err)
	}
	if err := m.To(s.to); err != nil {
		fmt.Printf("failed to set To address: %s", err)
	}

	m.Subject("task event")
	m.SetBodyString(mail.TypeTextPlain, message)

	c, err := mail.NewClient(s.host, s.port, mail.WithTLSPolicy(mail.NoTLS))
	if err != nil {
		fmt.Printf("failed to create mail client: %s", err)
	}

	if err = c.DialAndSend(m); err != nil {
		fmt.Printf("failed to send mail: %s", err)
	}
}
