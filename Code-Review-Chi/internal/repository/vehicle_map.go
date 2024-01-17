package repository

import (
	"app/internal"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// AddVehicle is a method that adds a vehicle to the repository
func (r *VehicleMap) AddVehicle(v internal.Vehicle) error {
	// verify if vehicle already exists in the repository
	if _, ok := r.db[v.Id]; ok {
		return internal.ErrVehicleAlreadyExists
	}

	// add vehicle to the repository
	r.db[v.Id] = v

	return nil
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// search vehicles by color and year
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrVehiclesNotFound
	}

	return
}

// FindByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
func (r *VehicleMap) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	fmt.Println("len of db:", len(r.db))
	// search vehicles by brand and year range
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrVehiclesNotFound
	}

	return
}

// GetAverageSpeedByBrand is a method that returns the average speed of vehicles by brand
func (r *VehicleMap) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	var totalSpeed float64
	var totalVehicles int

	// search vehicles by brand
	for _, value := range r.db {
		if value.Brand == brand {
			totalSpeed += value.MaxSpeed
			totalVehicles++
		}
	}

	if totalVehicles == 0 {
		return 0, internal.ErrVehiclesNotFound
	}

	averageSpeed = totalSpeed / float64(totalVehicles)

	return
}

// AddVehicles is a method that adds vehicles to the repository
func (r *VehicleMap) AddVehicles(v []internal.Vehicle) error {
	for _, vehicle := range v {
		err := r.AddVehicle(vehicle)
		if err != nil {
			return err
		}
	}

	return nil
}
