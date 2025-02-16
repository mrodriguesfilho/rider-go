package usecase

import (
	"fmt"
	"rider-go/internal/infra/database"

	"github.com/google/uuid"
)

type AcceptRide struct {
	accountRepository database.AccountRepository
	rideRepository    database.RideRepository
}

type AcceptRideInput struct {
	RideId   string
	DriverId string
}

type AcceptRideOutput struct {
	DriverId string
}

func NewAcceptRideUseCase(accountRepository database.AccountRepository, rideRepository database.RideRepository) *AcceptRide {
	return &AcceptRide{
		accountRepository: accountRepository,
		rideRepository:    rideRepository,
	}
}

func (a *AcceptRide) Execute(acceptRideInput AcceptRideInput) (AcceptRideOutput, error) {

	accountUuid, err := uuid.Parse(acceptRideInput.DriverId)

	if err != nil {
		return AcceptRideOutput{}, fmt.Errorf("cannot parse the received account id %s to uuid", acceptRideInput.DriverId)
	}

	driverAccount, err := a.accountRepository.GetById(accountUuid)

	if err != nil {
		return AcceptRideOutput{}, err
	}

	rideUuid, err := uuid.Parse(acceptRideInput.RideId)

	if err != nil {
		return AcceptRideOutput{}, fmt.Errorf("cannot parse the received ride id %s to uuid", acceptRideInput.RideId)
	}

	ride, err := a.rideRepository.GetById(rideUuid)

	if err != nil {
		return AcceptRideOutput{}, err
	}

	err = ride.AcceptRide(&driverAccount)

	if err != nil {
		return AcceptRideOutput{}, err
	}

	return AcceptRideOutput{
		DriverId: driverAccount.Id.String(),
	}, nil
}
