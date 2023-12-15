package repository

import (
	"app/internal/domain"
	"errors"
)

var (
	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")

	// ErrRepositoryDuplicatedId is returned when the id of a new vehicle is already registred
	ErrRepositoryDuplicatedId = errors.New("repository:  Duplicated ID")

	// ErrRepositoryDuplicatedId is returned when something happens weiting the file
	ErrRepositoryWritingStorage = errors.New("repository: error writing file")
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
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
