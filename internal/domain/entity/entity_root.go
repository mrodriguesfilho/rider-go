package entity

import (
	"rider-go/internal/domain/domainEvent"

	"github.com/google/uuid"
)

type EntityRoot struct {
	Id     uuid.UUID
	events []domainEvent.DomainEventInterface
}

func (r *EntityRoot) GetId() uuid.UUID {
	return r.Id
}

func (r *EntityRoot) RaiseEvent(event domainEvent.DomainEventInterface) {
	r.events = append(r.events, event)
}

func (r *EntityRoot) GetUncommitedEvents() []domainEvent.DomainEventInterface {

	var uncommitedEvents = make([]domainEvent.DomainEventInterface, 0)

	for _, e := range r.events {
		if !e.IsCommited() {
			uncommitedEvents = append(uncommitedEvents, e)
		}
	}

	return uncommitedEvents
}
