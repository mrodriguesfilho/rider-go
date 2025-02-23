package domainEvent

import "github.com/google/uuid"

type EventSource interface {
	GetId() uuid.UUID
	RaiseEvent(event DomainEventInterface)
	GetUncommitedEvents() []DomainEventInterface
}
