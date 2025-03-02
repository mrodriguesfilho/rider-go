package valueObjects

type GeoLocation struct {
	lat float64
	lon float64
}

func NewGeoLocation(lat float64, lon float64) GeoLocation {
	return GeoLocation{
		lat: lat,
		lon: lon,
	}
}

func (g GeoLocation) Equals(other GeoLocation) bool {
	return g.lat == other.lat && g.lon == other.lon
}

func (g GeoLocation) GetLat() float64 {
	return g.lat
}

func (g GeoLocation) GetLon() float64 {
	return g.lon
}
