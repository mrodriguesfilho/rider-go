package usecase

import (
	"rider-go/internal/application/event"
	eventhandlers "rider-go/internal/application/event/handlers"
	"rider-go/internal/domain/domainEvent"
	"rider-go/internal/domain/entity"
	"rider-go/internal/domain/valueObjects"
	inmemory "rider-go/internal/infra/database/InMemory"
	event_kafka_adapter "rider-go/internal/infra/event/kafka_adapter"
	"rider-go/internal/infra/payment"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type finishRideTestSuite struct {
	suite.Suite
	accountRepository  *inmemory.AccountRepositoryInMemory
	rideRepository     *inmemory.RideRepositoryInMemory
	requestRideUseCase RequestRideUseCase
	acceptRideUseCase  AcceptRideUseCase
	paymentService     payment.PaymentService
	eventDispatcher    event.EventDispatcher
	accountInMemoryDB  []entity.Account
	finishRideUseCase  FinishRideUseCase
}

func (f *finishRideTestSuite) SetupTest() {
	f.accountInMemoryDB = make([]entity.Account, 0)
	f.accountRepository = inmemory.NewAccountRepository(f.accountInMemoryDB)
	f.rideRepository = inmemory.NewRideRepositoryInMemory(make(map[uuid.UUID]entity.Ride))
	f.requestRideUseCase = *NewRequestRideUseCase(f.accountRepository, f.rideRepository)
	f.paymentService = payment.NewPaymentServiceInMemory()
	rideFinishedHandler := eventhandlers.NewRideFinishedEventHandler(f.rideRepository, f.accountRepository, f.paymentService)

	handlers := make(map[string]event.EventHandler)
	handlers[domainEvent.RideFinishedEventName] = rideFinishedHandler
	// eventChan := make(chan domainEvent.DomainEventInterface)
	// eventConsumer := event_in_memory_adapter.NewEventHandlerInMemory(eventChan, handlers)
	eventConsumer, _ := event_kafka_adapter.NewEventHandlerKafkaAdapter([]string{"localhost:9092"}, []string{"ride-domain-event-topic"}, handlers)
	eventConsumer.Listen()

	eventBroker, _ := event_kafka_adapter.NewEventBrokerKafkaAdapter([]string{"localhost:9092"})
	f.eventDispatcher = *event.NewEventDispatcher(eventBroker)
	f.acceptRideUseCase = *NewAcceptRideUseCase(f.accountRepository, f.rideRepository, f.eventDispatcher)
	f.finishRideUseCase = *NewFinishRideUseCase(f.rideRepository, f.accountRepository, f.eventDispatcher)
}

func (f *finishRideTestSuite) TestFinishRide1() {
	f.T().Run("It should allow a rider to finish a ride", func(t *testing.T) {
		passengerId, _ := uuid.NewUUID()
		passengerAccount := entity.Account{
			EntityRoot:  &entity.EntityRoot{Id: passengerId},
			Name:        "John Doe",
			Email:       "john.doe@gmail.com",
			Password:    "123",
			IsPassenger: true,
			IsDriver:    false,
		}

		driverAccountId, _ := uuid.NewUUID()
		driverAccount := entity.Account{
			EntityRoot:  &entity.EntityRoot{Id: driverAccountId},
			Name:        "Janes Doe",
			Email:       "janes.doe@gmail.com",
			Password:    "123",
			IsPassenger: false,
			IsDriver:    true,
		}

		f.accountRepository.Db = append(f.accountRepository.Db, passengerAccount)
		f.accountRepository.Db = append(f.accountRepository.Db, driverAccount)

		requestRideInput := RequestRideInput{
			PassengerId: passengerAccount.Id.String(),
			From:        valueObjects.NewGeoLocation(49, 45),
			To:          valueObjects.NewGeoLocation(50, 45),
		}

		requestRideOuput, _ := f.requestRideUseCase.Execute(requestRideInput)

		acceptRideInput := AcceptRideInput{
			RideId:   requestRideOuput.RideId.String(),
			DriverId: driverAccount.Id.String(),
		}

		f.acceptRideUseCase.Execute(acceptRideInput)

		finishRideInput := FinishRideInput{
			RideId:         requestRideOuput.RideId.String(),
			DriverLocation: valueObjects.NewGeoLocation(50, 45),
		}

		finishRideOuput, err := f.finishRideUseCase.Execute(finishRideInput)

		assert.Nil(t, err)
		assert.Equal(t, finishRideOuput.Status, entity.Finished)
	})
}

func TestFinishRideUseCase(t *testing.T) {
	suite.Run(t, new(finishRideTestSuite))
}
