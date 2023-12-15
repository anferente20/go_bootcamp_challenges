package service

import (
	"app/internal/domain"
	"errors"
)

var (
	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")
	// ErrServiceDuplicatedID is returned when vehicle id is already registered
	ErrServiceDuplicatedID = errors.New("service: Duplicated ID")
	// ErrRepositoryDuplicatedId is returned when something happens weiting the file
	ErrServiceWritingStorage = errors.New("service: error writing file")
)

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// FindAll returns all vehicles
	FindAll() (v []domain.Vehicle, err error)

	// Add a new Vehicle
	AddVehicle(vehicle domain.Vehicle) error

	// Add a new Vehicles
	AddVehicles(vehicles []domain.Vehicle) error

	// SearchByColorAndYear returns all vehicles filtered by color and year
	SearchByColorAndYear(color string, year int) (v []domain.Vehicle, err error)

	// SearchByBrandAndYears returns vehicles filtered by brand and year
	SearchByBrandAndYear(brand string, startYear, endYear int) (v []domain.Vehicle, err error)

	// SearchVehicleByFuelType returns a slice with all the vehicles with that fuel type
	SearchVehicleByFuelType(fuelType string) (v []domain.Vehicle, err error)

	// SearchVehicleByTransmissionType returns all vehicles with certain kind of transmission
	SearchVehicleByTransmissionType(transmissionType string) (v []domain.Vehicle, err error)

	// SearchVehicleByWeight returns all vehicles with certain kind of transmission
	SearchVehicleByWeight(minWeight, maxWeight float64) (v []domain.Vehicle, err error)

	// SearchVehiclesByDimensions returns cars with dimensions between range
	SearchVehiclesByDimensions(minHeight, maxHeight, minWidth, maxWidth float64) (v []domain.Vehicle, err error)

	// AverageSpeedByBrand returns a float with the average speed of some brand
	AverageSpeedByBrand(brand string) (speed float32, err error)

	// AverageCapacityByBrand returns a float with the average capacity of some brand
	AverageCapacityByBrand(brand string) (capacity int, err error)

	// UpdateVehicleSpeed allows to change a vehicle speed
	UpdateVehicleSpeed(vehicleId, newSpeed int) error

	// UpdateVehicleFuelType allows to change a vehicle fuel type
	UpdateVehicleFuelType(vehicleId int, newFuelType string) error

	// DeleteVehicle allows to delete an vehicle from db
	DeleteVehicle(vehicleId int) error
}
