package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRide(t *testing.T) {
	t.Run("It should create a ride", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, GeoLocation{Lat: 45, Lon: 45}, GeoLocation{Lat: 45, Lon: 46})

		assert.NotEqual(t, uuid.Nil, ride.Id)
		assert.Equal(t, ride.Status, Requested)
	})

	t.Run("It shouldn't allow to request a new ride", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		ride := NewRide(id, GeoLocation{Lat: 45, Lon: 45}, GeoLocation{Lat: 45, Lon: 46})

		assert.NotEqual(t, uuid.Nil, ride.Id)
		assert.Equal(t, ride.Status, Requested)
		assert.False(t, ride.StatusAllowedToRequestNewRide())
	})

	t.Run("It shouldn't allow an account without Driver flag true to accept a ride", func(t *testing.T) {

		id, _ := uuid.NewRandom()

		account := NewAccount("John Doe", "john.doe@gmail.com", "123", false, false)
		ride := NewRide(id, GeoLocation{Lat: 45, Lon: 45}, GeoLocation{Lat: 45, Lon: 46})

		err := ride.AcceptRide(account)

		assert.Error(t, err, "an account cannot accept a ride without drive flag marked as true")
	})
}
