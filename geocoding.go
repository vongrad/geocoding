package geocoding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const host = "https://maps.googleapis.com/maps/api/geocode/json"

// API for interaction with Google Maps Geocoding API
type API struct {
	// Google Maps API Key
	APIKey string
}

// Geocode returns the latitude and longitude given an address
func (api API) Geocode(address string) (float64, float64, error) {

	req, err := http.NewRequest("GET", host+"?address="+url.QueryEscape(address)+"&key="+api.APIKey, nil)

	if err != nil {
		return 0, 0, err
	}

	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		return 0, 0, err
	}

	var result Result

	dec := json.NewDecoder(res.Body)

	defer res.Body.Close()

	if err := dec.Decode(&result); err != nil {
		return 0, 0, err
	}

	if len(result.Results) == 0 {
		return 0, 0, fmt.Errorf("No results found matching %v", address)
	}

	lat := result.Results[0].Geometry.Location.Latitude
	lng := result.Results[0].Geometry.Location.Longitude

	return lat, lng, nil
}

// Result acquired from Google Maps Geocoding API
type Result struct {
	Results []Address `json:"results"`
}

// Address is a single geocoded address
type Address struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Geometry          AddressGeometry    `json:"geometry"`
	Types             []string           `json:"types"`
	PlaceID           string             `json:"place_id"`
}

// AddressComponent is a part of an address
type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// AddressGeometry is the location of a an address
type AddressGeometry struct {
	Location     LatLng       `json:"location"`
	LocationType string       `json:"location_type"`
	Bounds       LatLngBounds `json:"bounds"`
	Viewport     LatLngBounds `json:"viewport"`
	Types        []string     `json:"types"`
}

// LatLng is a tuple for latitude and longitude
type LatLng struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// LatLngBounds represents a bounded square area on the Earth.
type LatLngBounds struct {
	NorthEast LatLng `json:"northeast"`
	SouthWest LatLng `json:"southwest"`
}
