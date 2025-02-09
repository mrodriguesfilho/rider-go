package database

import (
	"fmt"
	"rider-go/internal/entity"

	"github.com/google/uuid"
)

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

func (r *RideRepositoryInMemory) GetLastRideByPassengerId(passengerId uuid.UUID) (entity.Ride, error) {
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
