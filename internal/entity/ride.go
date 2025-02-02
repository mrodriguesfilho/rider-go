package entity

import "github.com/google/uuid"

type RideStatus int

const (
	None RideStatus = iota
	Requested
	Completed
)

type Ride struct {
	Status      RideStatus
	PassengerId int
	Id          uuid.UUID
	From        GeoLocation
	To          GeoLocation
}

func NewRide(passengerId int, from GeoLocation, to GeoLocation) *Ride {
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
