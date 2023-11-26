package observer

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/y-nosuke/sample-task-api-go/app/notification/domain/event"
	tevent "github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
)

type SlackSubscriberImpl struct {
	token     string
	channelID string
	api       *slack.Client
}

func NewSlackSubscriberImpl(token string, channelID string) *SlackSubscriberImpl {
	api := slack.New(token)
	return &SlackSubscriberImpl{token, channelID, api}
}

func (s SlackSubscriberImpl) Subscribe(event event.DomainEvent) {
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

	channelID, timestamp, err := s.api.PostMessage(s.channelID, slack.MsgOptionText(message, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
