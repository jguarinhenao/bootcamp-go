package service

import (
	"app/internal"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

func validateVehicle(v *internal.Vehicle) (err error) {
	// Validar el ID del vehículo
	if v.Id <= 0 {
		return fmt.Errorf("%w: Id must be a positive integer", internal.ErrFieldRequired)
	}

	// Validar las dimensiones del vehículo
	if v.Height < 0 {
		return fmt.Errorf("%w: Height must be positive values", internal.ErrFieldRequired)
	}

	if v.Length < 0 {
		return fmt.Errorf("%w: Length must be positive values", internal.ErrFieldRequired)
	}

	if v.Width < 0 {
		return fmt.Errorf("%w: Width must be positive values", internal.ErrFieldRequired)
	}

	// Validar atributos adicionales del vehículo
	if err := validateVehicleAttributes(&v.VehicleAttributes); err != nil {
		return err
	}

	return nil
}

// validateVehicleAttributes is a function that validates the attributes of a vehicle
func validateVehicleAttributes(va *internal.VehicleAttributes) (err error) {

	if va.FabricationYear < 1900 {
		return fmt.Errorf("%w: FabricationYear must be 1900 or later", internal.ErrFieldRequired)
	}

	if va.Capacity <= 0 {
		return fmt.Errorf("%w: Capacity must be a positive integer", internal.ErrFieldRequired)
	}

	if va.MaxSpeed <= 0 {
		return fmt.Errorf("%w: MaxSpeed must be a positive value", internal.ErrFieldRequired)
	}

	if va.Weight <= 0 {
		return fmt.Errorf("%w: Weight must be a positive value", internal.ErrFieldRequired)
	}

	if va.Brand == "" {
		return fmt.Errorf("%w: Brand is required", internal.ErrFieldRequired)
	}

	if va.Model == "" {
		return fmt.Errorf("%w: Model is required", internal.ErrFieldRequired)
	}

	if va.Registration == "" {
		return fmt.Errorf("%w: Registration is required", internal.ErrFieldRequired)
	}

	if va.Color == "" {
		return fmt.Errorf("%w: Color is required", internal.ErrFieldRequired)
	}

	if va.FuelType == "" {
		return fmt.Errorf("%w: FuelType is required", internal.ErrFieldRequired)
	}

	if va.Transmission == "" {
		return fmt.Errorf("%w: Transmission is required", internal.ErrFieldRequired)
	}

	return nil
}

func validateYear(year int) (err error) {
	if year < 1900 || year > 2024 {
		return fmt.Errorf("%w: year must be between 1900 and 2024", internal.ErrFieldRequired)
	}

	return nil
}

func validateMaxSpeed(maxSpeed float64) (err error) {
	if maxSpeed <= 0 {
		return fmt.Errorf("%w: MaxSpeed must be a positive value", internal.ErrFieldRequired)
	}

	return nil
}

func validateWeightRanges(minRange float64, maxRange float64) (err error) {
	if minRange < 0 {
		return fmt.Errorf("%w: MinRange must be a positive value", internal.ErrFieldRequired)
	}

	if maxRange < 0 {
		return fmt.Errorf("%w: MaxRange must be a positive value", internal.ErrFieldRequired)
	}

	if minRange > maxRange {
		return fmt.Errorf("%w: MinRange must be less than MaxRange", internal.ErrFieldRequired)
	}

	return nil

}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// AddVehicle is a method that adds a vehicle to the repository
func (s *VehicleDefault) AddVehicle(v internal.Vehicle) (err error) {

	if err = validateVehicle(&v); err != nil {
		return err
	}

	err = s.rp.AddVehicle(v)
	if err != nil {
		switch err {
		case internal.ErrVehicleAlreadyExists:
			err = fmt.Errorf("%w: id", internal.ErrVehicleAlreadyExists)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}

	}

	return

}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {

	if err = validateYear(year); err != nil {
		return nil, err
	}

	v, err = s.rp.FindByColorAndYear(color, year)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: color %s and year %d", internal.ErrVehiclesNotFound, color, year)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {

	if err = validateYear(startYear); err != nil {
		return nil, fmt.Errorf("%w: startYear must be greater than 1900", internal.ErrFieldRequired)
	}

	if err = validateYear(endYear); err != nil {
		return nil, fmt.Errorf("%w: endYear must be less than 2024", internal.ErrFieldRequired)
	}

	v, err = s.rp.FindByBrandAndYearRange(brand, startYear, endYear)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: brand %s and year range %d - %d", internal.ErrVehiclesNotFound, brand, startYear, endYear)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.GetAverageSpeedByBrand(brand)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: brand %s", internal.ErrVehiclesNotFound, brand)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) AddVehicles(v []internal.Vehicle) (err error) {
	err = s.rp.AddVehicles(v)
	if err != nil {
		switch err {
		case internal.ErrVehicleAlreadyExists:
			err = fmt.Errorf("%w: id", internal.ErrVehicleAlreadyExists)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {

	v, err = s.rp.FindByFuelType(fuelType)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: fuelType %s", internal.ErrVehiclesNotFound, fuelType)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) DeleteVehicle(id int) (err error) {
	err = s.rp.DeleteVehicle(id)
	if err != nil {
		switch err {
		case internal.ErrVehicleNotFound:
			err = fmt.Errorf("%w: id %d", internal.ErrVehicleNotFound, id)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) FindByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {

	v, err = s.rp.FindByTransmissionType(transmissionType)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: transmissionType %s", internal.ErrVehiclesNotFound, transmissionType)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (r *VehicleDefault) UpdatePartials(id int, partials map[string]interface{}) (err error) {

	for key, value := range partials {
		switch key {
		case "max_speed":
			if err = validateMaxSpeed(value.(float64)); err != nil {
				return err
			}
		}
	}

	err = r.rp.UpdatePartials(id, partials)
	if err != nil {
		switch err {
		case internal.ErrVehicleNotFound:
			err = fmt.Errorf("%w: id %d", internal.ErrVehicleNotFound, id)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (s *VehicleDefault) GetAveragePassengersByBrand(brand string) (averagePassengers float64, err error) {
	averagePassengers, err = s.rp.GetAveragePassengersByBrand(brand)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: brand %s", internal.ErrVehiclesNotFound, brand)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return
}

func (r *VehicleDefault) FindByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {

	v, err = r.rp.FindByDimensions(minLength, maxLength, minWidth, maxWidth)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: dimensions %f - %f, %f - %f", internal.ErrVehiclesNotFound, minLength, maxLength, minWidth, maxWidth)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return

}

func (r *VehicleDefault) FindByWeightRange(minWeight float64, maxWeight float64) (v map[int]internal.Vehicle, err error) {

	if err = validateWeightRanges(minWeight, maxWeight); err != nil {
		return nil, err
	}

	v, err = r.rp.FindByWeightRange(minWeight, maxWeight)

	if err != nil {
		switch err {
		case internal.ErrVehiclesNotFound:
			err = fmt.Errorf("%w: weight %f - %f", internal.ErrVehiclesNotFound, minWeight, maxWeight)
		default:
			err = fmt.Errorf("%w", internal.ErrUnknown)
		}
	}

	return

}
