package database

import (
	"rider-go/internal/entity"

	"github.com/google/uuid"
)

type RideRepository interface {
	GetById(rideId uuid.UUID) (entity.Ride, error)
	GetLasRideByAccountId(passengerId uuid.UUID) (entity.Ride, error)
	Insert(ride *entity.Ride) error
}
