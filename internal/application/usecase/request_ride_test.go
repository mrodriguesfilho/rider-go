package usecase

import (
	"fmt"
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestRide(t *testing.T) {
	t.Run("It should request a ride", func(t *testing.T) {
		signUpInput := SignUpInput{
			Name:        "John Doe",
			Email:       "johndoe@gmail.com",
			Password:    "123123",
			IsPassenger: true,
			IsDriver:    false,
		}

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		signUpUseCase := NewSignUpUseCase(accountRepository)
		signUpOutput, errSignup := signUpUseCase.Execute(signUpInput)

		requestRideInput := RequestRideInput{
			PassengerId: signUpOutput.Id,
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
		requestRideOutput, _ := requestRideUseCase.Execute(requestRideInput)

		assert.Nil(t, errSignup)
		assert.NotEqual(t, uuid.Nil, requestRideOutput.RideId)
	})

	t.Run("It shouldn't request a ride for an account that doesn't have passenge flag as true", func(t *testing.T) {

		signUpInput := SignUpInput{
			Name:        "John Doe",
			Email:       "johndoe@gmail.com",
			Password:    "123123",
			IsPassenger: false,
			IsDriver:    false,
		}

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		signUpUseCase := NewSignUpUseCase(accountRepository)
		signUpOutput, errSignup := signUpUseCase.Execute(signUpInput)

		requestRideInput := RequestRideInput{
			PassengerId: signUpOutput.Id,
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
		requestRideOutput, errRequestRide := requestRideUseCase.Execute(requestRideInput)

		assert.Nil(t, errSignup)
		assert.Equal(t, "to request a ride the account has to have passenger flag marked as true", errRequestRide.Error())
		assert.Equal(t, uuid.Nil, requestRideOutput.RideId)
	})

	t.Run("It shouldn't request a ride for a passenger that has a ride with status different than completed", func(t *testing.T) {

		signUpInput := SignUpInput{
			Name:        "John Doe",
			Email:       "johndoe@gmail.com",
			Password:    "123123",
			IsPassenger: true,
			IsDriver:    false,
		}

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		signUpUseCase := NewSignUpUseCase(accountRepository)
		signUpOutput, errSignup := signUpUseCase.Execute(signUpInput)

		requestRideInput := RequestRideInput{
			PassengerId: signUpOutput.Id,
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
		requestRideOutputFirst, errRequestFirstRide := requestRideUseCase.Execute(requestRideInput)
		requestRideOutputSecond, errRequestSecondRide := requestRideUseCase.Execute(requestRideInput)

		assert.Nil(t, errSignup)
		assert.Nil(t, errRequestFirstRide)
		assert.NotEqual(t, uuid.Nil, requestRideOutputFirst.RideId)
		assert.Equal(t, "to request a ride the passenger's last ride must be completed", errRequestSecondRide.Error())
		assert.Equal(t, uuid.Nil, requestRideOutputSecond.RideId)
	})

	t.Run("It shouldn't request a ride for an account that doesn't exists", func(t *testing.T) {

		id, _ := uuid.NewRandom()
		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		requestRideInput := RequestRideInput{
			PassengerId: id.String(),
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
		requestRideOutputFirst, errRequestFirstRide := requestRideUseCase.Execute(requestRideInput)

		assert.NotNil(t, errRequestFirstRide)
		assert.Equal(t, uuid.Nil, requestRideOutputFirst.RideId)
		assert.Equal(t, fmt.Sprintf("no account with id %d was found", id), errRequestFirstRide.Error())
	})
}
