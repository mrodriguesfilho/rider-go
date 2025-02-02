package usecase

import (
	"fmt"
	"rider-go/internal/entity"
	"rider-go/internal/infra/database"

	"github.com/google/uuid"
)

type RequestRideUseCase struct {
	AccountRepository *database.AccountRepositoryInMemory
	RideRepository    database.RideRepository
}

type RequestRideInput struct {
	PassengerId int
	From        entity.GeoLocation
	To          entity.GeoLocation
}

type RequestRideOutput struct {
	RideId uuid.UUID
}

func NewRequestRideUseCase(accountRepository *database.AccountRepositoryInMemory, rideRepository database.RideRepository) *RequestRideUseCase {
	return &RequestRideUseCase{
		AccountRepository: accountRepository,
		RideRepository:    rideRepository,
	}
}

func (r *RequestRideUseCase) Execute(requestRideInput RequestRideInput) (RequestRideOutput, error) {

	passenger, err := r.AccountRepository.GetById(requestRideInput.PassengerId)

	if err != nil {
		return RequestRideOutput{}, err
	}

	if !passenger.IsPassenger {
		return RequestRideOutput{}, fmt.Errorf("to request a ride the account has to have passenger flag marked as true")
	}

	lastRide, err := r.RideRepository.GetLastRideByPassengerId(passenger.Id)

	if err != nil {
		return RequestRideOutput{}, err
	}

	if lastRide.Id != uuid.Nil && !lastRide.StatusAllowedToRequestNewRide() {
		return RequestRideOutput{}, fmt.Errorf("to request a ride the passenger's last ride must be completed")
	}

	newRide := entity.NewRide(requestRideInput.PassengerId, requestRideInput.From, requestRideInput.To)

	err = r.RideRepository.Insert(newRide)

	if err != nil {
		return RequestRideOutput{}, err
	}

	return RequestRideOutput{
		newRide.Id,
	}, nil
}
