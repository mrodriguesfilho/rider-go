package event_handlers

import (
	"encoding/json"
	"rider-go/internal/domain/domainEvent"
	"rider-go/internal/infra/database/repository"
	"rider-go/internal/infra/payment"
)

type RideAcceptedHandler struct {
	paymentService    payment.PaymentService
	rideRepository    repository.RideRepository
	accountRepository repository.AccountRepository
}

func NewRideAcceptedEventHandler(rideRepository repository.RideRepository, accountRepository repository.AccountRepository, paymentService payment.PaymentService) *RideAcceptedHandler {
	return &RideAcceptedHandler{
		rideRepository:    rideRepository,
		accountRepository: accountRepository,
		paymentService:    paymentService,
	}
}

func (r *RideAcceptedHandler) Handle(domainEventInterface domainEvent.DomainEventInterface) {

	var rideAcceptedEvent domainEvent.RideAccepted
	json.Unmarshal([]byte(domainEventInterface.GetPayload()), &rideAcceptedEvent)

	ride, err := r.rideRepository.GetById(rideAcceptedEvent.RideId)

	if err != nil {
		return
	}

	passenger, err := r.accountRepository.GetById(ride.PassengerId)

	if err != nil {
		return
	}

	r.paymentService.Debit(passenger.Email, ride.Fare)
}
