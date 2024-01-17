package internal

import "errors"

var (
	ErrFieldRequired = errors.New("field is required")

	ErrInvalidFieldEnum = errors.New("invalid field enum")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	AddVehicle(v Vehicle) (err error)

	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)

	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)

	AddVehicles(v []Vehicle) (err error)
}
