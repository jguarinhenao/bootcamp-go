package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

type UpdateSpeedJSON struct {
	MaxSpeed float64 `json:"max_speed"`
}

type UpdateFuelJSON struct {
	FuelType string `json:"fuel_type"`
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

type DimensionQueryParams struct {
	LengthRange string `query:"length"`
	WidthRange  string `query:"width"`
}

func parseRange(rangeStr string) (min float64, max float64, err error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		err = fmt.Errorf("range format is invalid: %s", rangeStr)
		return
	}

	min, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}

	max, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return
	}

	if min < 0 || max < 0 || min > max {
		err = errors.New("range values are invalid")
	}

	return
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

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: id must be an integer")
			return
		}

		var body UpdateSpeedJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		if err := h.sv.UpdatePartials(idInt, map[string]interface{}{"max_speed": body.MaxSpeed}); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
		})

	}
}

func (h *VehicleDefault) FindByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fuelType := chi.URLParam(r, "type")

		if fuelType == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		v, err := h.sv.FindByFuelType(fuelType)
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

func (h *VehicleDefault) DeleteVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: id must be an integer")
			return
		}

		if err := h.sv.DeleteVehicle(idInt); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}

			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{})

	}
}

func (h *VehicleDefault) FindByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		transmissionType := chi.URLParam(r, "type")

		if transmissionType == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		v, err := h.sv.FindByTransmissionType(transmissionType)
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

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: id must be an integer")
			return
		}

		var body UpdateFuelJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		if err := h.sv.UpdatePartials(idInt, map[string]interface{}{"fuel_type": body.FuelType}); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
		})

	}
}

func (h *VehicleDefault) GetAveragePassengersByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")

		if brand == "" {
			response.Text(w, http.StatusBadRequest, "invalid query params")
			return
		}

		v, err := h.sv.GetAveragePassengersByBrand(brand)
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

		averageCapacity := map[string]any{
			"average_capacity": v,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageCapacity,
		})

	}
}

func (h *VehicleDefault) FindByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var queryParams DimensionQueryParams
		if err := r.ParseForm(); err != nil {
			response.Text(w, http.StatusBadRequest, "error parsing query parameters")
			return
		}

		queryParams.LengthRange = r.Form.Get("length")
		queryParams.WidthRange = r.Form.Get("width")

		minLength, maxLength, err := parseRange(queryParams.LengthRange)
		if err != nil {
			response.Text(w, http.StatusBadRequest, err.Error())
			return
		}

		minWidth, maxWidth, err := parseRange(queryParams.WidthRange)
		if err != nil {
			response.Text(w, http.StatusBadRequest, err.Error())
			return
		}

		v, err := h.sv.FindByDimensions(minLength, maxLength, minWidth, maxWidth)
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

func (h *VehicleDefault) FindByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		minWeight, err := strconv.ParseFloat(r.URL.Query().Get("min"), 64)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: min must be a float")
			return
		}

		maxWeight, err := strconv.ParseFloat(r.URL.Query().Get("max"), 64)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid query params: max must be a float")
			return
		}

		v, err := h.sv.FindByWeightRange(minWeight, maxWeight)
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
