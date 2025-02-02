package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRide(t *testing.T) {
	t.Run("It should create a ride", func(t *testing.T) {
		ride := NewRide(1, GeoLocation{Lat: 45, Lon: 45}, GeoLocation{Lat: 45, Lon: 46})

		assert.NotEqual(t, uuid.Nil, ride.Id)
		assert.Equal(t, ride.Status, Requested)
	})

	t.Run("It shouldn't allow to request a new ride", func(t *testing.T) {
		ride := NewRide(1, GeoLocation{Lat: 45, Lon: 45}, GeoLocation{Lat: 45, Lon: 46})

		assert.NotEqual(t, uuid.Nil, ride.Id)
		assert.Equal(t, ride.Status, Requested)
		assert.False(t, ride.StatusAllowedToRequestNewRide())
	})
}
