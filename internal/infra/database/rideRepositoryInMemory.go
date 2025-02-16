package database

import (
	"fmt"
	"rider-go/internal/entity"

	"github.com/google/uuid"
)

type RideRepositoryInMemory struct {
	db map[uuid.UUID]entity.Ride
}

func (r *RideRepositoryInMemory) GetById(id uuid.UUID) (entity.Ride, error) {

	ride, exists := r.db[id]

	if !exists {
		return entity.Ride{}, fmt.Errorf("no ride was found with id %s", id.String())
	}

	return ride, nil
}

func (r *RideRepositoryInMemory) Insert(ride *entity.Ride) error {

	if _, exists := r.db[ride.Id]; exists {
		return fmt.Errorf("there is already a ride with ID %s on the database", ride.Id)
	}

	r.db[ride.Id] = *ride

	return nil
}

func (r *RideRepositoryInMemory) GetLasRideByAccountId(passengerId uuid.UUID) (entity.Ride, error) {
	for _, ride := range r.db {
		if ride.PassengerId == passengerId {
			return ride, nil
		}
	}

	return entity.Ride{}, nil
}

func NewRideRepository(db map[uuid.UUID]entity.Ride) RideRepository {
	return &RideRepositoryInMemory{
		db: db,
	}
}
