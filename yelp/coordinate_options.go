package yelp

import (
	"errors"
	"fmt"
)

// CoordinateOptions are used with complex searches for locations.
// The geographic coordinate format is defined as:
// ll=latitude,longitude,accuracy,altitude,altitude_accuracy
type CoordinateOptions struct {
	Latitude         Float // Latitude of geo-point to search near (required)
	Longitude        Float // Longitude of geo-point to search near (required)
	Accuracy         Float // Accuracy of latitude, longitude (optional)
	Altitude         Float // Altitude (optional)
	AltitudeAccuracy Float // Accuracy of altitude (optional)
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
// that match the defined values.
func (o CoordinateOptions) getParameters() (params map[string]string, err error) {
	// coordinate requires at least a latitude and longitude - others are option
	if !o.Latitude.Valid || !o.Longitude.Valid {
		return nil, errors.New("latitude and longitude are required fields for a coordinate based search")
	}

	ll := fmt.Sprintf("%v,%v", o.Latitude.Float64, o.Longitude.Float64)
	if o.Accuracy.Valid {
		ll += fmt.Sprintf(",%v", o.Accuracy.Float64)
	}
	if o.Altitude.Valid {
		ll += fmt.Sprintf(",%v", o.Altitude.Float64)
	}
	if o.AltitudeAccuracy.Valid {
		ll += fmt.Sprintf(",%v", o.AltitudeAccuracy.Float64)
	}

	return map[string]string{
		"ll": ll,
	}, nil
}
