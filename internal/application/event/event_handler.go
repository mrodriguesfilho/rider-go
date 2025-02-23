package event

import (
	"rider-go/internal/domain/domainEvent"
)

type EventHandler interface {
	Handle(event domainEvent.DomainEventInterface)
}
