package internal

import "errors"

var (
	// ErrVehicleAlreadyExists is an error that represents a vehicle already exists in the repository
	ErrVehicleAlreadyExists = errors.New("vehicle already exists in the repository")
	ErrUnmarshal            = errors.New("error unmarshaling file")
	ErrMarshal              = errors.New("error marshaling file")
	ErrWriteFile            = errors.New("error writing file")
	ErrUnknown              = errors.New("unknown error")
	ErrVehiclesNotFound     = errors.New("vehicles not found")
	ErrVehicleNotFound      = errors.New("vehicle not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	AddVehicle(v Vehicle) (err error)

	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)

	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)

	AddVehicles(v []Vehicle) (err error)

	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)
	DeleteVehicle(id int) (err error)
	FindByTransmissionType(transmissionType string) (v map[int]Vehicle, err error)
	UpdatePartials(id int, partials map[string]interface{}) (err error)

	GetAveragePassengersByBrand(brand string) (averagePassengers float64, err error)

	FindByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v map[int]Vehicle, err error)

	FindByWeightRange(minWeight float64, maxWeight float64) (v map[int]Vehicle, err error)
}
