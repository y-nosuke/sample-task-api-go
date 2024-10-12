package repository

import (
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/event"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/database/dao"
	"golang.org/x/xerrors"
)

var (
	TaskCreated     []byte
	TaskEdited      []byte
	TaskCompleted   []byte
	TaskUnCompleted []byte
	TaskDeleted     []byte
)

func init() {
	var eventTypeID uuid.UUID
	eventTypeID, _ = uuid.Parse("019271db-fcae-7760-b8cf-d5a773e02b2c")
	TaskCreated, _ = eventTypeID.MarshalBinary()

	eventTypeID, _ = uuid.Parse("019271db-fcae-797a-ad72-6849999ee5ee")
	TaskEdited, _ = eventTypeID.MarshalBinary()

	eventTypeID, _ = uuid.Parse("019271db-fcae-7864-b7d9-b61db784aba4")
	TaskCompleted, _ = eventTypeID.MarshalBinary()

	eventTypeID, _ = uuid.Parse("019271db-fcae-7bbd-9557-d468a7f149cc")
	TaskUnCompleted, _ = eventTypeID.MarshalBinary()

	eventTypeID, _ = uuid.Parse("019271db-fcae-7bda-8ed4-fe79343cf736")
	TaskDeleted, _ = eventTypeID.MarshalBinary()

}
func ETaskEvent(taskEvent *event.TaskEvent, eventTypeID []byte, registerBy uuid.UUID) (*dao.ETaskEvent, error) {
	id, err := taskEvent.ID().MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("taskEvent.ID().MarshalBinary(): %w", err)
	}

	taskID, err := taskEvent.TaskID.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("taskEvent.TaskID().MarshalBinary(): %w", err)
	}

	byteOccurredBy, err := taskEvent.OccurredBy.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("occurredBy.MarshalBinary(): %w", err)
	}

	byteRegisterBy, err := registerBy.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("registerBy.MarshalBinary(): %w", err)
	}

	eTaskEvent := &dao.ETaskEvent{
		ID:          id,
		EventTypeID: eventTypeID,
		TaskID:      taskID,
		OccurredBy:  byteOccurredBy,
		OccurredAt:  taskEvent.OccurredAt,
		RegisterBy:  byteRegisterBy,
	}

	return eTaskEvent, nil
}

func ETaskCreated(taskCreated *event.TaskCreated, userId uuid.UUID) (*dao.ETaskEvent, *dao.ETaskCreated, error) {
	eTaskEvent, err := ETaskEvent(&taskCreated.TaskEvent, TaskCreated, userId)
	if err != nil {
		return nil, nil, err
	}

	eTaskCreated := &dao.ETaskCreated{
		EventID:    eTaskEvent.ID,
		Title:      taskCreated.Title,
		Detail:     null.StringFromPtr(taskCreated.Detail),
		Completed:  taskCreated.Completed,
		Deadline:   null.TimeFromPtr(taskCreated.Deadline),
		RegisterBy: eTaskEvent.RegisterBy,
	}

	return eTaskEvent, eTaskCreated, nil
}

func ETaskUpdated(taskUpdated *event.TaskEdited, userId uuid.UUID) (*dao.ETaskEvent, *dao.ETaskEdited, error) {
	eTaskEvent, err := ETaskEvent(&taskUpdated.TaskEvent, TaskEdited, userId)
	if err != nil {
		return nil, nil, err
	}

	byteUserId, err := userId.MarshalBinary()
	if err != nil {
		return nil, nil, xerrors.Errorf("userId.MarshalBinary(): %w", err)
	}

	eTaskCreated := &dao.ETaskEdited{
		EventID:    eTaskEvent.ID,
		Title:      taskUpdated.Title,
		Detail:     null.StringFromPtr(taskUpdated.Detail),
		Completed:  taskUpdated.Completed,
		Deadline:   null.TimeFromPtr(taskUpdated.Deadline),
		RegisterBy: byteUserId,
	}

	return eTaskEvent, eTaskCreated, nil
}

func ETaskCompleted(taskCompleted *event.TaskCompleted, userId uuid.UUID) (*dao.ETaskEvent, error) {
	eTaskEvent, err := ETaskEvent(&taskCompleted.TaskEvent, TaskCompleted, userId)
	if err != nil {
		return nil, err
	}

	return eTaskEvent, nil
}

func ETaskUnCompleted(taskUnCompleted *event.TaskUnCompleted, userId uuid.UUID) (*dao.ETaskEvent, error) {
	eTaskEvent, err := ETaskEvent(&taskUnCompleted.TaskEvent, TaskUnCompleted, userId)
	if err != nil {
		return nil, err
	}

	return eTaskEvent, nil
}

func ETaskDeleted(taskDeleted *event.TaskDeleted, userId uuid.UUID) (*dao.ETaskEvent, error) {
	eTaskEvent, err := ETaskEvent(&taskDeleted.TaskEvent, TaskDeleted, userId)
	if err != nil {
		return nil, err
	}

	return eTaskEvent, nil
}
