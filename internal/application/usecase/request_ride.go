package usecase

import (
	"fmt"
	"rider-go/internal/domain/entity"
	"rider-go/internal/infra/database/repository"

	"github.com/google/uuid"
)

type RequestRideUseCase struct {
	accountRepository repository.AccountRepository
	rideRepository    repository.RideRepository
}

type RequestRideInput struct {
	PassengerId string
	From        entity.GeoLocation
	To          entity.GeoLocation
}

type RequestRideOutput struct {
	RideId uuid.UUID
}

func NewRequestRideUseCase(accountRepository repository.AccountRepository, rideRepository repository.RideRepository) *RequestRideUseCase {
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

	if lastRide.EntityRoot != nil && lastRide.Id != uuid.Nil && !lastRide.StatusAllowedToRequestNewRide() {
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
