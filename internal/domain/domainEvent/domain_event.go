package domainEvent

import (
	"encoding/json"

	"github.com/google/uuid"
)

type DomainEventInterface interface {
	GetEventName() string
	IsCommited() bool
	GetId() uuid.UUID
	GetPayload() string
	MarkAsCommited()
}

type DomainEventType struct {
	Id        uuid.UUID
	Commited  bool
	EventName string
	Payload   string
}

func NewDomainEvent(name string, payload any) *DomainEventType {

	payloadString, _ := json.Marshal(payload)

	return &DomainEventType{
		Id:        uuid.New(),
		EventName: name,
		Payload:   string(payloadString),
	}
}

func (r *DomainEventType) IsCommited() bool {
	return r.Commited
}

func (r *DomainEventType) GetEventName() string {
	return r.EventName
}

func (r *DomainEventType) GetId() uuid.UUID {
	return r.Id
}

func (r *DomainEventType) GetPayload() string {
	return r.Payload
}

func (r *DomainEventType) MarkAsCommited() {
	r.Commited = true
}
