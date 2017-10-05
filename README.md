# Google Maps Geocoding API for GoLang

Simple Google Maps Geocoding API that translate address into latitude & longitude

## Install

```
go get github.com/vongrad/geocoding
```

## Usage

``` go
api := API{apiKey: "API_KEY"}

lat, lng, err := api.Geocode(address)
```
