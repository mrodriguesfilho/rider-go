package event

import (
	"rider-go/internal/domain/domainEvent"

	"github.com/google/uuid"
)

type EventDispatcher struct {
	eventSourceMap map[uuid.UUID]domainEvent.EventSource
	eventBroker    EventBroker
}

func NewEventDispatcher(eventBroker EventBroker) *EventDispatcher {
	return &EventDispatcher{
		eventSourceMap: map[uuid.UUID]domainEvent.EventSource{},
		eventBroker:    eventBroker,
	}
}

func (e *EventDispatcher) Add(eventSource domainEvent.EventSource) {
	e.eventSourceMap[eventSource.GetId()] = eventSource
}

func (e *EventDispatcher) Commit() {
	for _, eventSource := range e.eventSourceMap {
		events := eventSource.GetUncommitedEvents()
		for _, event := range events {
			if !event.IsCommited() {
				e.eventBroker.Publish(event)
				event.MarkAsCommited()
			}
		}
	}
}
