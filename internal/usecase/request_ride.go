package usecase

import (
	"fmt"
	"rider-go/internal/entity"
	"rider-go/internal/infra/database"

	"github.com/google/uuid"
)

type RequestRideUseCase struct {
	accountRepository database.AccountRepository
	rideRepository    database.RideRepository
}

type RequestRideInput struct {
	PassengerId string
	From        entity.GeoLocation
	To          entity.GeoLocation
}

type RequestRideOutput struct {
	RideId uuid.UUID
}

func NewRequestRideUseCase(accountRepository database.AccountRepository, rideRepository database.RideRepository) *RequestRideUseCase {
	return &RequestRideUseCase{
		accountRepository: accountRepository,
		rideRepository:    rideRepository,
	}
}

func (r *RequestRideUseCase) Execute(requestRideInput RequestRideInput) (RequestRideOutput, error) {

	idParsed, err := uuid.Parse(requestRideInput.PassengerId)

	if err != nil {
		return RequestRideOutput{}, fmt.Errorf("id was in invalid format. execpected uuid got %s", requestRideInput.PassengerId)
	}

	passenger, err := r.accountRepository.GetById(idParsed)

	if err != nil {
		return RequestRideOutput{}, err
	}

	if !passenger.IsPassenger {
		return RequestRideOutput{}, fmt.Errorf("to request a ride the account has to have passenger flag marked as true")
	}

	lastRide, err := r.rideRepository.GetLasRideByAccountId(idParsed)

	if err != nil {
		return RequestRideOutput{}, err
	}

	if lastRide.Id != uuid.Nil && !lastRide.StatusAllowedToRequestNewRide() {
		return RequestRideOutput{}, fmt.Errorf("to request a ride the passenger's last ride must be completed")
	}

	newRide := entity.NewRide(passenger.Id, requestRideInput.From, requestRideInput.To)

	err = r.rideRepository.Insert(newRide)

	if err != nil {
		return RequestRideOutput{}, err
	}

	return RequestRideOutput{
		newRide.Id,
	}, nil
}
