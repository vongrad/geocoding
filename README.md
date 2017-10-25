[![Build Status](https://travis-ci.org/vongrad/geocoding.svg?branch=master)](https://travis-ci.org/vongrad/geocoding)

# Google Maps Geocoding API for GoLang

Simple Google Maps Geocoding API that translate address into latitude & longitude

## Install

```
go get github.com/vongrad/geocoding
```

## Usage

``` go
api := API{APIKey: "API_KEY"}

lat, lng, err := api.Geocode(address)
```
