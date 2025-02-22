package usecase

import (
	"fmt"
	"rider-go/internal/domain/entity"
	"rider-go/internal/infra/database/repository"
	"sync"

	"github.com/google/uuid"
)

type AcceptRide struct {
	accountRepository repository.AccountRepository
	rideRepository    repository.RideRepository
}

type AcceptRideInput struct {
	RideId   string
	DriverId string
}

type AcceptRideOutput struct {
	DriverId string
}

func NewAcceptRideUseCase(accountRepository repository.AccountRepository, rideRepository repository.RideRepository) *AcceptRide {
	return &AcceptRide{
		accountRepository: accountRepository,
		rideRepository:    rideRepository,
	}
}

func (a *AcceptRide) Execute(acceptRideInput AcceptRideInput) (AcceptRideOutput, error) {

	accountUuid, parseErr := uuid.Parse(acceptRideInput.DriverId)

	if parseErr != nil {
		return AcceptRideOutput{}, fmt.Errorf("cannot parse the received account id %s to uuid", acceptRideInput.DriverId)
	}

	rideUuid, parseErr := uuid.Parse(acceptRideInput.RideId)

	if parseErr != nil {
		return AcceptRideOutput{}, fmt.Errorf("cannot parse the received ride id %s to uuid", acceptRideInput.RideId)
	}

	var wg sync.WaitGroup
	var driverAccount entity.Account
	var ride entity.Ride
	var driverAccountErr, rideErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		driverAccount, driverAccountErr = a.accountRepository.GetById(accountUuid)
	}()

	go func() {
		defer wg.Done()
		ride, rideErr = a.rideRepository.GetById(rideUuid)
	}()

	wg.Wait()

	if driverAccountErr != nil {
		return AcceptRideOutput{}, driverAccountErr
	}

	if rideErr != nil {
		return AcceptRideOutput{}, rideErr
	}

	acceptRideErr := ride.AcceptRide(driverAccount)

	if acceptRideErr != nil {
		return AcceptRideOutput{}, acceptRideErr
	}

	return AcceptRideOutput{
		DriverId: ride.DriverId.String(),
	}, nil
}
