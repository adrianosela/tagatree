package objects

import (
	"errors"
	"math"
)

var (
	errLocationTypeInvalid             = errors.New("only location type \"Point\" is supported")
	errLocationCoordinatesInvalid      = errors.New("invalid coordinates array")
	errLocationCoordinatesBadLongitude = errors.New("longitude must respect x: -180.0 < x < 180.0")
	errLocationCoordinatesBadLatitude  = errors.New("latitude must respect y: -90.0 < y < 90.0")
)

const (
	typePoint = "Point"
)

// Location is a GeoJSON type.
type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

// NewLocation returns a GeoJSON Point with longitude and latitude.
func NewLocation(long, lat float64) Location {
	return Location{Type: typePoint, Coordinates: []float64{long, lat}}
}

// Validate validates a location
func (l *Location) Validate() error {
	if l.Type == "" { // allow
		l.Type = typePoint
	}
	if l.Type != typePoint {
		return errLocationTypeInvalid
	}
	if len(l.Coordinates) != 2 {
		return errLocationCoordinatesInvalid
	}
	if math.Abs(l.Coordinates[0]) > 180.0 {
		return errLocationCoordinatesBadLongitude
	}
	if math.Abs(l.Coordinates[1]) > 90.0 {
		return errLocationCoordinatesBadLatitude
	}
	return nil
}
