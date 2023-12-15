package service

import (
	"app/internal/domain"
	"app/internal/vehicles/repository"
	"errors"
	"fmt"
)

// NewDefault returns a new instance of a vehicle service.
func NewDefault(rp repository.RepositoryVehicle) *Default {
	return &Default{rp: rp}
}

// Default is an struct that represents a vehicle service.
type Default struct {
	rp repository.RepositoryVehicle
}

// FindAll returns all vehicles.
func (s *Default) FindAll() (v []domain.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindAll()
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) SearchByColorAndYear(color string, year int) (v []domain.Vehicle, err error) {
	// get vehicles filtered by color and year
	v, err = s.rp.SearchByColorAndYear(color, year)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) SearchVehicleByFuelType(fuelType string) (v []domain.Vehicle, err error) {
	// get vehicles  by fuel type
	v, err = s.rp.SearchVehicleByFuelType(fuelType)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) SearchVehicleByTransmissionType(transmissionType string) (v []domain.Vehicle, err error) {
	// get vehicles  by fuel type
	v, err = s.rp.SearchVehicleByTransmissionType(transmissionType)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) SearchByBrandAndYear(brand string, startYear, endYear int) (v []domain.Vehicle, err error) {
	// get vehicles filtered brand and year
	v, err = s.rp.SearchByBrandAndYear(brand, startYear, endYear)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) SearchVehiclesByDimensions(minHeight, maxHeight, minWidth, maxWidth float64) (v []domain.Vehicle, err error) {
	// get vehicles filtered brand and year
	v, err = s.rp.SearchVehiclesByDimensions(minHeight, maxHeight, minWidth, maxWidth)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) SearchVehicleByWeight(minWeight, maxWeight float64) (v []domain.Vehicle, err error) {
	// get vehicles filtered brand and year
	v, err = s.rp.SearchVehicleByWeight(minWeight, maxWeight)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) AverageSpeedByBrand(brand string) (speed float32, err error) {
	// get vehicles filtered brand and year
	speed, err = s.rp.AverageSpeedByBrand(brand)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) AverageCapacityByBrand(brand string) (capacity int, err error) {
	// get vehicles filtered brand and year
	capacity, err = s.rp.AverageCapacityByBrand(brand)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) AddVehicle(vehicle domain.Vehicle) error {

	// add new vehicle
	err := s.rp.AddVehicle(vehicle)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryDuplicatedId) {
			err = fmt.Errorf("%w. %v", ErrServiceDuplicatedID, err)
			return err
		}
		if errors.Is(err, repository.ErrRepositoryWritingStorage) {
			err = fmt.Errorf("%w. %v", ErrServiceWritingStorage, err)
			return err
		}
		return err
	}

	return nil

}

func (s *Default) AddVehicles(vehicle []domain.Vehicle) error {

	// add new vehicle
	err := s.rp.AddVehicles(vehicle)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryDuplicatedId) {
			err = fmt.Errorf("%w. %v", ErrServiceDuplicatedID, err)
			return err
		}
		if errors.Is(err, repository.ErrRepositoryWritingStorage) {
			err = fmt.Errorf("%w. %v", ErrServiceWritingStorage, err)
			return err
		}
		return err
	}

	return nil

}

func (s *Default) UpdateVehicleSpeed(vehicleId, newSpeed int) error {
	// update vehicle speed
	err := s.rp.UpdateVehicleSpeed(vehicleId, newSpeed)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
		}
		if errors.Is(err, repository.ErrRepositoryWritingStorage) {
			err = fmt.Errorf("%w. %v", ErrServiceWritingStorage, err)
			return err
		}
	}
	return err
}

func (s *Default) UpdateVehicleFuelType(vehicleId int, newFuelType string) error {
	// update vehicle speed
	err := s.rp.UpdateVehicleFuelType(vehicleId, newFuelType)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
		}
		if errors.Is(err, repository.ErrRepositoryWritingStorage) {
			err = fmt.Errorf("%w. %v", ErrServiceWritingStorage, err)
			return err
		}
	}
	return err
}

func (s *Default) DeleteVehicle(vehicleId int) error {
	// update vehicle speed
	err := s.rp.DeleteVehicle(vehicleId)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
		}
		if errors.Is(err, repository.ErrRepositoryWritingStorage) {
			err = fmt.Errorf("%w. %v", ErrServiceWritingStorage, err)
			return err
		}
	}
	return err
}
