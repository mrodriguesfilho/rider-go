package eventAdapters

import (
	"fmt"
	"rider-go/internal/application/event"
	"rider-go/internal/domain/domainEvent"
)

type EventhHandlerInMemoryAdapter struct {
	eventChan chan domainEvent.DomainEventInterface
	handlers  map[string]event.EventHandler
}

func NewEventHandlerInMemory(eventChan chan domainEvent.DomainEventInterface, handlers map[string]event.EventHandler) *EventhHandlerInMemoryAdapter {
	eventHandlerInMemoryAdapter := EventhHandlerInMemoryAdapter{
		eventChan: eventChan,
		handlers:  handlers,
	}
	return &eventHandlerInMemoryAdapter
}

func (eh *EventhHandlerInMemoryAdapter) Listen() {
	go func() {
		for event := range eh.eventChan {
			if handler, exists := eh.handlers[event.EventName()]; exists {
				handler.Handle(event)
			} else {
				fmt.Printf("No handler found for the given event")
			}
		}
	}()
}
