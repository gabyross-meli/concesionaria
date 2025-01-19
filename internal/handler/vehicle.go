package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bootcamp-go/web/response"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
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
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
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
		// read the body
		body := r.Body
		vehicle := models.VehicleDoc{}
		err := json.NewDecoder(body).Decode(&vehicle)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = h.sv.AddVehicle(vehicle)
		if err != nil {
			if strings.Contains(err.Error(), "400") {
				response.JSON(w, http.StatusConflict, err.Error())
			} else if strings.Contains(err.Error(), "409") {
				response.JSON(w, http.StatusBadRequest, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
		} else {
			response.JSON(w, http.StatusCreated, "201 Created: Vehículo creado exitosamente")
		}
	}
}
