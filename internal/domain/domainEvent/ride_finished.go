package domainEvent

import "github.com/google/uuid"

type RideFinished struct {
	RideId uuid.UUID
}

const RideFinishedEventName string = "RideFinished"

func NewRideFinishedEvent(rideId uuid.UUID) *DomainEventType {
	return NewDomainEvent(RideFinishedEventName, RideFinished{RideId: rideId})
}
