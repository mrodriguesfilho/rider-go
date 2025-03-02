package event_in_memory_adapter

import (
	"rider-go/internal/domain/domainEvent"
)

type EventBrokerInMemoryAdapter struct {
	eventChan chan domainEvent.DomainEventInterface
}

func NewEventBrokerInMemory(eventChan chan domainEvent.DomainEventInterface) *EventBrokerInMemoryAdapter {
	return &EventBrokerInMemoryAdapter{
		eventChan: eventChan,
	}
}

func (i *EventBrokerInMemoryAdapter) Publish(event domainEvent.DomainEventInterface) {
	i.eventChan <- event
}
