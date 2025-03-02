package entity

import (
	"rider-go/internal/domain/valueObjects"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRide(t *testing.T) {
	t.Run("It should create a ride", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))

		assert.NotEqual(t, uuid.Nil, ride.Id)
		assert.Equal(t, ride.Status, Requested)
	})

	t.Run("It shouldn't allow to request a new ride", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))

		assert.NotEqual(t, uuid.Nil, ride.Id)
		assert.Equal(t, ride.Status, Requested)
		assert.False(t, ride.StatusAllowedToRequestNewRide())
	})

	t.Run("It shouldn't allow an account without Driver flag true to accept a ride", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		account := NewAccount("John Doe", "john.doe@gmail.com", "123", false, false)
		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))

		err := ride.AcceptRide(*account)

		assert.Error(t, err, "an account cannot accept a ride without drive flag marked as true")
	})

	t.Run("It should allow a ride to be finished", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))
		account := NewAccount("John Doe", "john.doe@gmail.com", "123", false, true)
		ride.AcceptRide(*account)
		err := ride.FinishRide(valueObjects.NewGeoLocation(45, 46))

		assert.Nil(t, err)
		assert.Equal(t, ride.Status, Finished)
	})

	t.Run("It shouldn't allow a ride to be finished twice", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))
		account := NewAccount("John Doe", "john.doe@gmail.com", "123", false, true)
		ride.AcceptRide(*account)
		_ = ride.FinishRide(valueObjects.NewGeoLocation(45, 46))
		err := ride.FinishRide(valueObjects.NewGeoLocation(45, 46))

		assert.Error(t, err)
		assert.EqualError(t, err, "this ride was already finished")
	})

	t.Run("It shouldn't allow a ride to be finished before reaches the destination", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))
		account := NewAccount("John Doe", "john.doe@gmail.com", "123", false, true)
		ride.AcceptRide(*account)
		err := ride.FinishRide(valueObjects.NewGeoLocation(45, 45.5))

		assert.Equal(t, ride.Status, Accepted)
		assert.Error(t, err)
		assert.EqualError(t, err, "the driver location must be the same as the ride destination")
	})

	t.Run("It shouldn't allow a ride to be finished if it wasn't accepted", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, valueObjects.NewGeoLocation(45, 45), valueObjects.NewGeoLocation(45, 46))
		err := ride.FinishRide(valueObjects.NewGeoLocation(45, 46))

		assert.Error(t, err)
		assert.Equal(t, ride.Status, Requested)
		assert.EqualError(t, err, "a ride cannot be finished without being accepted")
	})
}
