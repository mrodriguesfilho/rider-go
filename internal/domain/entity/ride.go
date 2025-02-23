package entity

import (
	"errors"
	"rider-go/internal/domain/domainEvent"
	"rider-go/internal/domain/valueObjects"

	"github.com/google/uuid"
)

type RideStatus int

const (
	None RideStatus = iota
	Requested
	Accepted
	Completed
)

type Ride struct {
	*EntityRoot
	Status      RideStatus
	PassengerId uuid.UUID
	DriverId    uuid.UUID
	From        GeoLocation
	To          GeoLocation
	Fare        valueObjects.Money
	DriverFare  valueObjects.Money
}

func NewRide(passengerId uuid.UUID, from GeoLocation, to GeoLocation) *Ride {

	return &Ride{
		EntityRoot:  &EntityRoot{Id: uuid.New()},
		PassengerId: passengerId,
		From:        from,
		To:          to,
		Status:      Requested,
		Fare:        valueObjects.NewMoney(10, valueObjects.USD),
		DriverFare:  valueObjects.NewMoney(8, valueObjects.USD),
	}
}

func (r *Ride) StatusAllowedToRequestNewRide() bool {
	return r.Status == Completed
}

func (r *Ride) AcceptRide(driverAccount Account) error {
	if !driverAccount.IsDriver {
		return errors.New("an account cannot accept a ride without driver flag marked as true")
	}

	if r.Status != Requested {
		return errors.New("a ride cannot be accepted outside of requested status")
	}

	r.DriverId = driverAccount.Id
	r.Status = Accepted

	r.RaiseEvent(domainEvent.NewRideAcceptedEvent(r.Id))

	return nil
}
