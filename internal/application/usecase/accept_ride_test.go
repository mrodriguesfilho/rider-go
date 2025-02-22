package usecase

import (
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAcceptRideUseCase(t *testing.T) {
	t.Run("It should allow a rider to accept a ride", func(t *testing.T) {

		passengerId, _ := uuid.NewUUID()
		passengerAccount := entity.Account{
			Id:          passengerId,
			Name:        "John Doe",
			Email:       "john.doe@gmail.com",
			Password:    "123",
			IsPassenger: true,
			IsDriver:    false,
		}

		driverAccountId, _ := uuid.NewUUID()
		driverAccount := entity.Account{
			Id:          driverAccountId,
			Name:        "Janes Doe",
			Email:       "janes.doe@gmail.com",
			Password:    "123",
			IsPassenger: false,
			IsDriver:    true,
		}

		inMemoryDatabase := make([]entity.Account, 0)
		inMemoryDatabase = append(inMemoryDatabase, passengerAccount)
		inMemoryDatabase = append(inMemoryDatabase, driverAccount)
		accountRepository := inmemory.NewAccountRepository(inMemoryDatabase)

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

		rideRepository := inmemory.NewRideRepository(make(map[uuid.UUID]entity.Ride))
		requestRideUseCase := NewRequestRideUseCase(accountRepository, rideRepository)
		requestRideOuput, _ := requestRideUseCase.Execute(requestRideInput)

		acceptRideInput := AcceptRideInput{
			RideId:   requestRideOuput.RideId.String(),
			DriverId: driverAccount.Id.String(),
		}

		acceptRideUseCase := NewAcceptRideUseCase(accountRepository, rideRepository)
		acceptRideOuput, _ := acceptRideUseCase.Execute(acceptRideInput)

		assert.NotNil(t, uuid.Nil, acceptRideOuput.DriverId)
		assert.Equal(t, acceptRideOuput.DriverId, driverAccount.Id.String())
	})

	t.Run("It shouldn't allow an account to acept a ride if it isn't a driver", func(t *testing.T) {

		passengerId, _ := uuid.NewUUID()
		passengerAccount := entity.Account{
			Id:          passengerId,
			Name:        "John Doe",
			Email:       "john.doe@gmail.com",
			Password:    "123",
			IsPassenger: true,
			IsDriver:    false,
		}

		driverAccountId, _ := uuid.NewUUID()
		driverAccount := entity.Account{
			Id:          driverAccountId,
			Name:        "Janes Doe",
			Email:       "janes.doe@gmail.com",
			Password:    "123",
			IsPassenger: true,
			IsDriver:    false,
		}

		inMemoryDatabase := make([]entity.Account, 0)
		inMemoryDatabase = append(inMemoryDatabase, passengerAccount)
		inMemoryDatabase = append(inMemoryDatabase, driverAccount)
		accountRepository := inmemory.NewAccountRepository(inMemoryDatabase)

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

		rideRepository := inmemory.NewRideRepository(make(map[uuid.UUID]entity.Ride))
		requestRideUseCase := NewRequestRideUseCase(accountRepository, rideRepository)
		requestRideOuput, _ := requestRideUseCase.Execute(requestRideInput)

		acceptRideInput := AcceptRideInput{
			RideId:   requestRideOuput.RideId.String(),
			DriverId: driverAccount.Id.String(),
		}

		acceptRideUseCase := NewAcceptRideUseCase(accountRepository, rideRepository)
		acceptRideOuput, err := acceptRideUseCase.Execute(acceptRideInput)

		assert.Equal(t, "", acceptRideOuput.DriverId)
		assert.Equal(t, "an account cannot accept a ride without driver flag marked as true", err.Error())
	})
}
