package geocoding

import "testing"

const (
	address   string  = "1600 Amphitheatre Parkway, Mountain View, CA"
	latitude  float64 = 37.4216548
	longitude float64 = -122.0856374
)

func TestGeocode(t *testing.T) {

	api := API{APIKey: ""}

	lat, lng, err := api.Geocode(address)

	if err != nil {
		t.Errorf("1600 Amphitheatre Parkway, Mountain View, CA: Expected error to be nil, ~ Received %v", err)
	}

	if lat != latitude || lng != longitude {
		t.Errorf("1600 Amphitheatre Parkway, Mountain View, CA: Expected (%f, %f), ~ Received (%f, %f)", latitude, longitude, lat, lng)
	}

}
