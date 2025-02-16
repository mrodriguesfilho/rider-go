package entity

import (
	"errors"

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
	Id          uuid.UUID
	Status      RideStatus
	PassengerId uuid.UUID
	DriverId    uuid.UUID
	From        GeoLocation
	To          GeoLocation
}

func NewRide(passengerId uuid.UUID, from GeoLocation, to GeoLocation) *Ride {
	return &Ride{
		Id:          uuid.New(),
		PassengerId: passengerId,
		From:        from,
		To:          to,
		Status:      Requested,
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

	return nil
}
