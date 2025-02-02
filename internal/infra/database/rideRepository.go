package database

import (
	"fmt"
	"rider-go/internal/entity"

	"github.com/google/uuid"
)

type RideRepository interface {
	GetLastRideByPassengerId(passengerId int) (entity.Ride, error)
	Insert(ride *entity.Ride) error
}

type RideRepositoryInMemory struct {
	Db map[uuid.UUID]entity.Ride
}

func (r *RideRepositoryInMemory) Insert(ride *entity.Ride) error {

	if _, exists := r.Db[ride.Id]; exists {
		return fmt.Errorf("there is already a ride with ID %s on the database", ride.Id)
	}

	r.Db[ride.Id] = *ride

	return nil
}

func (r *RideRepositoryInMemory) GetLastRideByPassengerId(passengerId int) (entity.Ride, error) {
	for _, ride := range r.Db {
		if ride.PassengerId == passengerId {
			return ride, nil
		}
	}

	return entity.Ride{}, nil
}

func NewRideRepository(db map[uuid.UUID]entity.Ride) RideRepository {
	return &RideRepositoryInMemory{
		Db: db,
	}
}
