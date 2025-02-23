package domainEvent

import "github.com/google/uuid"

type RideAccepted struct {
	RideId uuid.UUID
}

const RideAccepetedEventName string = "RideAccepted"

func NewRideAcceptedEvent(rideId uuid.UUID) *DomainEventType {
	return NewDomainEvent(RideAccepetedEventName, RideAccepted{RideId: rideId})
}
