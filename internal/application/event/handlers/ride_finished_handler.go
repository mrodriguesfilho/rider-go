package event_handlers

import (
	"encoding/json"
	"rider-go/internal/domain/domainEvent"
	"rider-go/internal/infra/database/repository"
	"rider-go/internal/infra/payment"
)

type RideFinishedHandler struct {
	paymentService    payment.PaymentService
	rideRepository    repository.RideRepository
	accountRepository repository.AccountRepository
}

func NewRideFinishedEventHandler(rideRepository repository.RideRepository, accountRepository repository.AccountRepository, paymentService payment.PaymentService) *RideFinishedHandler {
	return &RideFinishedHandler{
		rideRepository:    rideRepository,
		accountRepository: accountRepository,
		paymentService:    paymentService,
	}
}

func (r *RideFinishedHandler) Handle(domainEventInterface domainEvent.DomainEventInterface) {

	var rideFinishedEvent domainEvent.RideFinished
	json.Unmarshal([]byte(domainEventInterface.GetPayload()), &rideFinishedEvent)

	ride, err := r.rideRepository.GetById(rideFinishedEvent.RideId)

	if err != nil {
		return
	}

	driver, err := r.accountRepository.GetById(ride.DriverId)

	if err != nil {
		return
	}

	r.paymentService.Credit(driver.Email, ride.DriverFare)
}
