package usecase

import (
	"rider-go/internal/application/event"
	"rider-go/internal/domain/entity"
	"rider-go/internal/domain/valueObjects"
	"rider-go/internal/infra/database/repository"

	"github.com/google/uuid"
)

type FinishRideInput struct {
	RideId         string
	DriverLocation valueObjects.GeoLocation
}

type FinishRideOutput struct {
	RideId string
	Status entity.RideStatus
}

type FinishRideUseCase struct {
	rideRepository    repository.RideRepository
	accountRepository repository.AccountRepository
	eventDispatcher   event.EventDispatcher
}

func NewFinishRideUseCase(rideRepository repository.RideRepository, accountRepository repository.AccountRepository, eventDispatcher event.EventDispatcher) *FinishRideUseCase {
	return &FinishRideUseCase{
		rideRepository:    rideRepository,
		accountRepository: accountRepository,
		eventDispatcher:   eventDispatcher,
	}
}

func (f *FinishRideUseCase) Execute(finishRideInput FinishRideInput) (FinishRideOutput, error) {

	rideId, err := uuid.Parse(finishRideInput.RideId)

	if err != nil {
		return FinishRideOutput{}, err
	}

	ride, err := f.rideRepository.GetById(rideId)

	if err != nil {
		return FinishRideOutput{}, err
	}

	err = ride.FinishRide(finishRideInput.DriverLocation)

	if err != nil {
		return FinishRideOutput{}, err
	}

	err = f.rideRepository.Update(ride)

	if err != nil {
		return FinishRideOutput{}, err
	}

	f.eventDispatcher.Add(ride)
	f.eventDispatcher.Commit()

	return FinishRideOutput{
		RideId: ride.Id.String(),
		Status: ride.Status,
	}, nil
}
