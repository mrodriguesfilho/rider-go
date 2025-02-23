package usecase

import (
	"rider-go/internal/application/event"
	"rider-go/internal/application/eventhandlers"
	"rider-go/internal/domain/domainEvent"
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	"rider-go/internal/infra/eventAdapters"
	"rider-go/internal/infra/payment"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type accountRideTestSuite struct {
	suite.Suite
	accountRepository  *inmemory.AccountRepositoryInMemory
	rideRepository     *inmemory.RideRepositoryInMemory
	requestRideUseCase RequestRideUseCase
	acceptRideUseCase  AcceptRideUseCase
	paymentService     payment.PaymentService
	eventDispatcher    event.EventDispatcher
	accountInMemoryDB  []entity.Account
}

func (s *accountRideTestSuite) SetupTest() {
	s.accountInMemoryDB = make([]entity.Account, 0)
	s.accountRepository = inmemory.NewAccountRepository(s.accountInMemoryDB)
	s.rideRepository = inmemory.NewRideRepositoryInMemory(make(map[uuid.UUID]entity.Ride))
	s.requestRideUseCase = *NewRequestRideUseCase(s.accountRepository, s.rideRepository)
	s.paymentService = payment.NewPaymentServiceInMemory()
	rideAccepetedHandler := eventhandlers.NewRideAcceptedEventHandler(s.rideRepository, s.accountRepository, s.paymentService)

	handlers := make(map[string]event.EventHandler)
	handlers[domainEvent.RideAccepetedEventName] = rideAccepetedHandler
	eventChan := make(chan domainEvent.DomainEventInterface)
	eventConsumer := eventAdapters.NewEventHandlerInMemory(eventChan, handlers)
	eventConsumer.Listen()

	eventBroker := eventAdapters.NewEventBrokerInMemory(eventChan)
	s.eventDispatcher = *event.NewEventDispatcher(eventBroker)
	s.acceptRideUseCase = *NewAcceptRideUseCase(s.accountRepository, s.rideRepository, s.eventDispatcher)
}

func (s *accountRideTestSuite) TestAcceptRide1() {
	s.T().Run("It should allow a rider to accept a ride", func(t *testing.T) {
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

		s.accountRepository.Db = append(s.accountRepository.Db, passengerAccount)
		s.accountRepository.Db = append(s.accountRepository.Db, driverAccount)

		requestRideInput := RequestRideInput{
			PassengerId: passengerAccount.Id.String(),
			From: entity.GeoLocation{
				Lat: 49,
				Lon: 45,
			},
			To: entity.GeoLocation{
				Lat: 50,
				Lon: 45,
			},
		}

		requestRideOuput, _ := s.requestRideUseCase.Execute(requestRideInput)

		acceptRideInput := AcceptRideInput{
			RideId:   requestRideOuput.RideId.String(),
			DriverId: driverAccount.Id.String(),
		}

		acceptRideOuput, _ := s.acceptRideUseCase.Execute(acceptRideInput)

		assert.Equal(t, acceptRideOuput.DriverId, driverAccount.Id.String())
	})
}

func (s *accountRideTestSuite) TestAcceptRide2() {

	s.T().Run("It shouldn't allow an account to acept a ride if it isn't a driver", func(t *testing.T) {
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
			IsPassenger: true,
			IsDriver:    false,
		}

		s.accountRepository.Db = append(s.accountRepository.Db, passengerAccount)
		s.accountRepository.Db = append(s.accountRepository.Db, driverAccount)

		requestRideInput := RequestRideInput{
			PassengerId: passengerAccount.Id.String(),
			From: entity.GeoLocation{
				Lat: 49,
				Lon: 45,
			},
			To: entity.GeoLocation{
				Lat: 50,
				Lon: 45,
			},
		}

		requestRideOuput, _ := s.requestRideUseCase.Execute(requestRideInput)

		acceptRideInput := AcceptRideInput{
			RideId:   requestRideOuput.RideId.String(),
			DriverId: driverAccount.Id.String(),
		}

		acceptRideOuput, err := s.acceptRideUseCase.Execute(acceptRideInput)

		assert.Equal(t, "", acceptRideOuput.DriverId)
		assert.Equal(t, "an account cannot accept a ride without driver flag marked as true", err.Error())
	})
}

func TestAcceptRideUseCase(t *testing.T) {
	suite.Run(t, new(accountRideTestSuite))
}
