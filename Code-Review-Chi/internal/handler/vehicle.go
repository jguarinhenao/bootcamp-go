package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) AddVehicle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var body VehicleJSON

		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		vehicle := internal.Vehicle{
			Id: body.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		if err := h.sv.AddVehicle(vehicle); err != nil {
			fmt.Printf("error: %v", err)
			switch {
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.Text(w, http.StatusConflict, "vehicle already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.VehicleAttributes.Brand,
			Model:           vehicle.VehicleAttributes.Model,
			Registration:    vehicle.VehicleAttributes.Registration,
			Color:           vehicle.VehicleAttributes.Color,
			FabricationYear: vehicle.VehicleAttributes.FabricationYear,
			Capacity:        vehicle.VehicleAttributes.Capacity,
			MaxSpeed:        vehicle.VehicleAttributes.MaxSpeed,
			FuelType:        vehicle.VehicleAttributes.FuelType,
			Transmission:    vehicle.VehicleAttributes.Transmission,
			Height:          vehicle.VehicleAttributes.Dimensions.Height,
			Length:          vehicle.VehicleAttributes.Dimensions.Length,
			Width:           vehicle.VehicleAttributes.Dimensions.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehicle added successfully",
			"data":    data,
		})

	}

}

func (h *VehicleDefault) FindByColorAndYear() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")

		if color == "" || year == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		yearInt, err := strconv.Atoi(year)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: year must be an integer")
			return
		}

		v, err := h.sv.FindByColorAndYear(color, yearInt)
		if err != nil {

			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}

			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) FindByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")
		startYear := chi.URLParam(r, "start_year")
		endYear := chi.URLParam(r, "end_year")

		if brand == "" || startYear == "" || endYear == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		startYearInt, err := strconv.Atoi(startYear)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: start_year must be an integer")
			return
		}

		endYearInt, err := strconv.Atoi(endYear)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: end_year must be an integer")
			return
		}

		v, err := h.sv.FindByBrandAndYearRange(brand, startYearInt, endYearInt)
		if err != nil {

			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}

			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}

}

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")

		if brand == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		v, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {

			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}

			return
		}

		averageSpeed := map[string]any{
			"average_speed": v,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed,
		})

	}
}

func (h *VehicleDefault) AddVehicles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body []VehicleJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		vehicles := make([]internal.Vehicle, len(body))
		for key, value := range body {
			vehicles[key] = internal.Vehicle{
				Id: value.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           value.Brand,
					Model:           value.Model,
					Registration:    value.Registration,
					Color:           value.Color,
					FabricationYear: value.FabricationYear,
					Capacity:        value.Capacity,
					MaxSpeed:        value.MaxSpeed,
					FuelType:        value.FuelType,
					Transmission:    value.Transmission,
					Dimensions: internal.Dimensions{
						Height: value.Height,
						Length: value.Length,
						Width:  value.Width,
					},
				},
			}
		}

		if err := h.sv.AddVehicles(vehicles); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.Text(w, http.StatusConflict, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}

			return
		}
	}
}
