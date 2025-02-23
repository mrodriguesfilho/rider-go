package domainEvent

import (
	"github.com/google/uuid"
)

type DomainEventInterface interface {
	EventName() string
	IsCommited() bool
	GetId() uuid.UUID
	GetPayload() interface{}
	MarkAsCommited()
}

type DomainEventType struct {
	id        uuid.UUID
	commited  bool
	eventName string
	payload   interface{}
}

func NewDomainEvent(name string, payload interface{}) *DomainEventType {

	return &DomainEventType{
		id:        uuid.New(),
		eventName: name,
		payload:   payload,
	}
}

func (r *DomainEventType) IsCommited() bool {
	return r.commited
}

func (r *DomainEventType) EventName() string {
	return r.eventName
}

func (r *DomainEventType) GetId() uuid.UUID {
	return r.id
}

func (r *DomainEventType) GetPayload() interface{} {
	return r.payload
}

func (r *DomainEventType) MarkAsCommited() {
	r.commited = true
}
