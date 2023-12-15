package handler

import (
	"app/internal/domain"
	"app/internal/vehicles/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// VehicleJSON is an struct that represents a vehicle in json format.
type VehicleJSON struct {
	ID           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

// NewVehicleDefault returns a new instance of a vehicle handler.
func NewVehicleDefault(sv service.ServiceVehicle) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is an struct that contains handlers for vehicle.
type VehicleDefault struct {
	sv service.ServiceVehicle
}

// GetAll returns all vehicles.
func (c *VehicleDefault) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.FindAll()
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) SearchByColorAndYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}

		color := ctx.Param("color")

		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.SearchByColorAndYear(color, year)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) SearchVehicleByFuelType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		fuelType := ctx.Param("type")

		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.SearchVehicleByFuelType(fuelType)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) SearchVehicleByTransmissionType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		transmission := ctx.Param("type")

		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.SearchVehicleByTransmissionType(transmission)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) SearchByBrandAndYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		startYear, err := strconv.Atoi(ctx.Param("start_year"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}
		endYear, err := strconv.Atoi(ctx.Param("end_year"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}

		brand := ctx.Param("brand")

		// process
		// - get vehicles from the service
		vehicles, err := c.sv.SearchByBrandAndYear(brand, startYear, endYear)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) SearchVehicleByWeight() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		minWeight, err := strconv.ParseFloat(ctx.Query("min"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}
		maxWeight, err := strconv.ParseFloat(ctx.Query("max"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}

		// process
		// - get vehicles from the service
		vehicles, err := c.sv.SearchVehicleByWeight(minWeight, maxWeight)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) SearchVehiclesByDimensions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		height := strings.Split(ctx.Query("height"), "-")

		minHeight, err := strconv.ParseFloat(height[0], 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}
		maxHeight, err := strconv.ParseFloat(height[1], 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}

		width := strings.Split(ctx.Query("width"), "-")

		minWidth, err := strconv.ParseFloat(width[0], 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}
		maxWidth, err := strconv.ParseFloat(width[1], 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos ingresads mal formados o incompletos"})
			return
		}

		fmt.Println(minHeight, maxHeight, minWidth, maxWidth)
		// process
		// - get vehicles from the service
		vehicles, err := c.sv.SearchVehiclesByDimensions(minHeight, maxHeight, minWidth, maxWidth)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}
func (c *VehicleDefault) AverageSpeedByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		brand := ctx.Param("brand")

		// process
		// - get vehicles from the service
		speed, err := c.sv.AverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := map[string]interface{}{
			"average_speed": speed,
			"brand":         brand,
		}

		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) AverageCapacityByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		brand := ctx.Param("brand")

		// process
		// - get vehicles from the service
		speed, err := c.sv.AverageCapacityByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := map[string]interface{}{
			"average_capacity": speed,
			"brand":            brand,
		}

		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) AddVehicle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req domain.Vehicle
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos del vehículo mal formados o incompletos"})
			return
		}
		fmt.Println(req)
		err := c.sv.AddVehicle(req)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceDuplicatedID):
				ctx.JSON(http.StatusConflict, map[string]any{"message": "Indentificador del vehículo ya existente"})
			case errors.Is(err, service.ErrServiceWritingStorage):
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "error writing database"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles

		ctx.JSON(http.StatusCreated, map[string]any{"message": "success to add a vehicles", "data": req})
	}
}

func (c *VehicleDefault) AddVehicles() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req []domain.Vehicle
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos del vehículo mal formados o incompletos"})
			return
		}

		err := c.sv.AddVehicles(req)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceDuplicatedID):
				ctx.JSON(http.StatusConflict, map[string]any{"message": "Indentificador del vehículo ya existente"})
			case errors.Is(err, service.ErrServiceWritingStorage):
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "error writing database"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// responseß
		// - serialize vehicles

		ctx.JSON(http.StatusCreated, map[string]any{"message": "success to add a vehicles", "data": req})
	}
}

func (c *VehicleDefault) UpdateVehicleSpeed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var req map[string]int
		if err := ctx.ShouldBind(&req); err != nil || req["max_speed"] < 0 {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Velocidad mal formada o fuera de rango"})
			return
		}

		vehicleId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Id mal formado o fuera de rango"})
			return
		}
		// process
		// - get vehicles from the service
		err = c.sv.UpdateVehicleSpeed(vehicleId, req["max_speed"])
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			case errors.Is(err, service.ErrServiceWritingStorage):
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "error writing database"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		ctx.JSON(http.StatusOK, map[string]any{"message": "Velocidad del vehículo actualizada exitosamente."})
	}
}

func (c *VehicleDefault) UpdateVehicleFuelType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var req map[string]string
		if err := ctx.ShouldBind(&req); err != nil || req["fuel_type"] == "" {
			fmt.Print(err)
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Velocidad mal formada o fuera de rango"})
			return
		}

		vehicleId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Id mal formado o fuera de rango"})
			return
		}
		// process
		// - get vehicles from the service
		err = c.sv.UpdateVehicleFuelType(vehicleId, req["fuel_type"])
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			case errors.Is(err, service.ErrServiceWritingStorage):
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "error writing database"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		ctx.JSON(http.StatusOK, map[string]any{"message": "Velocidad del vehículo actualizada exitosamente."})
	}
}

func (c *VehicleDefault) DeleteVehicle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		vehicleId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Id mal formado o fuera de rango"})
			return
		}
		// process
		// - get vehicles from the service
		err = c.sv.DeleteVehicle(vehicleId)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			case errors.Is(err, service.ErrServiceWritingStorage):
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "error writing database"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		ctx.JSON(http.StatusNoContent, map[string]any{"message": "Vehículo eliminado exitosamente."})
	}
}
