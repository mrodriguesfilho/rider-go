package event

import "rider-go/internal/domain/domainEvent"

type EventBroker interface {
	Publish(event domainEvent.DomainEventInterface)
}
