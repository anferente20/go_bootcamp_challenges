package repository

import (
	"app/internal/domain"
	"app/pkg/loader"
	"fmt"
	"slices"
)

// NewVehicleSlice returns a new instance of a vehicle repository in an slice.
func NewVehicleSlice(ld loader.VehicleJSON) *VehicleSlice {
	data, err := ld.Load()
	if err != nil {
		return &VehicleSlice{}
	}

	return &VehicleSlice{
		db:     data.Data,
		lastId: data.LastId,
		loader: ld,
	}
}

// VehicleSlice is an struct that represents a vehicle repository in an slice.
type VehicleSlice struct {
	// db is the database of vehicles.
	db []domain.Vehicle
	// lastId is the last id of the database.
	lastId int
	// loader that allows modify json file
	loader loader.VehicleJSON
}

// FindAll returns all vehicles
func (s *VehicleSlice) FindAll() (v []domain.Vehicle, err error) {
	fmt.Println(s.db)
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	// make a copy of the database
	v = make([]domain.Vehicle, len(s.db))
	copy(v, s.db)
	return
}

func (s *VehicleSlice) SearchVehicleByTransmissionType(transmissionType string) (v []domain.Vehicle, err error) {
	vehiclesFiltered := []domain.Vehicle{}
	for _, v := range s.db {
		if v.Attributes.Transmission == transmissionType {
			vehiclesFiltered = append(vehiclesFiltered, v)
		}
	}

	if len(vehiclesFiltered) == 0 {
		err = ErrRepositoryVehicleNotFound
		return []domain.Vehicle{}, err
	}
	return vehiclesFiltered, nil
}

func (s *VehicleSlice) SearchByColorAndYear(color string, year int) (v []domain.Vehicle, err error) {
	vehiclesFiltered := []domain.Vehicle{}
	for _, v := range s.db {
		if v.Attributes.Color == color && v.Attributes.Year == year {
			vehiclesFiltered = append(vehiclesFiltered, v)
		}
	}

	if len(vehiclesFiltered) == 0 {
		err = ErrRepositoryVehicleNotFound
		return []domain.Vehicle{}, err
	}
	return vehiclesFiltered, nil
}

func (s *VehicleSlice) SearchByBrandAndYear(brand string, startYear, endYear int) (v []domain.Vehicle, err error) {
	vehiclesFiltered := []domain.Vehicle{}
	for _, v := range s.db {
		if v.Attributes.Brand == brand && v.Attributes.Year >= startYear && v.Attributes.Year <= endYear {
			vehiclesFiltered = append(vehiclesFiltered, v)
		}
	}

	if len(vehiclesFiltered) == 0 {
		err = ErrRepositoryVehicleNotFound
		return []domain.Vehicle{}, err
	}
	return vehiclesFiltered, nil
}

func (s *VehicleSlice) SearchVehiclesByDimensions(minHeight, maxHeight, minWidth, maxWidth float64) (v []domain.Vehicle, err error) {
	vehiclesFiltered := []domain.Vehicle{}
	for _, v := range s.db {
		if v.Attributes.Height >= minHeight && v.Attributes.Height <= maxHeight && v.Attributes.Width >= minWidth && v.Attributes.Width <= maxWidth {
			vehiclesFiltered = append(vehiclesFiltered, v)
		}
	}

	if len(vehiclesFiltered) == 0 {
		err = ErrRepositoryVehicleNotFound
		return []domain.Vehicle{}, err
	}
	return vehiclesFiltered, nil
}

func (s *VehicleSlice) AverageSpeedByBrand(brand string) (speed float32, err error) {
	var averageSpeed float32
	var countCars float32
	for _, v := range s.db {
		if v.Attributes.Brand == brand {
			averageSpeed += float32(v.Attributes.MaxSpeed)
			countCars++
		}
	}

	if countCars == 0 {
		err = ErrRepositoryVehicleNotFound
		return 0, err
	}
	return averageSpeed / countCars, nil
}

func (s *VehicleSlice) AverageCapacityByBrand(brand string) (capacity int, err error) {
	var averageCapacity int
	var countCars int
	for _, v := range s.db {
		if v.Attributes.Brand == brand {
			averageCapacity += v.Attributes.Passengers
			countCars++
		}
	}

	if countCars == 0 {
		err = ErrRepositoryVehicleNotFound
		return 0, err
	}
	return averageCapacity / countCars, nil
}

func (s *VehicleSlice) SearchVehicleByFuelType(fuelType string) (v []domain.Vehicle, err error) {
	vehiclesFiltered := []domain.Vehicle{}
	for _, v := range s.db {
		if v.Attributes.FuelType == fuelType {
			vehiclesFiltered = append(vehiclesFiltered, v)
		}
	}

	if len(vehiclesFiltered) == 0 {
		err = ErrRepositoryVehicleNotFound
		return []domain.Vehicle{}, err
	}
	return vehiclesFiltered, nil
}

func (s *VehicleSlice) SearchVehicleByWeight(minWeight, maxWeight float64) (v []domain.Vehicle, err error) {
	vehiclesFiltered := []domain.Vehicle{}
	for _, v := range s.db {
		if minWeight < v.Attributes.Weight && v.Attributes.Weight < maxWeight {
			vehiclesFiltered = append(vehiclesFiltered, v)
		}
	}

	if len(vehiclesFiltered) == 0 {
		err = ErrRepositoryVehicleNotFound
		return []domain.Vehicle{}, err
	}
	return vehiclesFiltered, nil
}

func (s *VehicleSlice) AddVehicle(vehicle domain.Vehicle) error {

	alreadyRegistered := false
	for _, v := range s.db {
		if v.ID == vehicle.ID {
			alreadyRegistered = true
			break
		}
	}
	if alreadyRegistered {
		return ErrRepositoryDuplicatedId
	}
	s.db = append(s.db, vehicle)
	s.lastId = len(s.db)
	err := s.loader.Write(s.db)
	if err != nil {
		return ErrRepositoryWritingStorage
	}
	return nil
}

func (s *VehicleSlice) AddVehicles(vehicles []domain.Vehicle) error {
	alreadyRegistered := false
	for _, v1 := range vehicles {
		for _, v2 := range s.db {
			if v1.ID == v2.ID {
				alreadyRegistered = true
				break
			}
		}
	}
	if alreadyRegistered {
		return ErrRepositoryDuplicatedId
	}

	s.db = append(s.db, vehicles...)
	err := s.loader.Write(s.db)
	if err != nil {
		return ErrRepositoryWritingStorage
	}
	return nil
}

func (s *VehicleSlice) UpdateVehicleSpeed(vehicleId, newSpeed int) error {
	found := false
	for _, v := range s.db {
		if v.ID == vehicleId {
			index := slices.Index(s.db, v)
			v.Attributes.MaxSpeed = newSpeed

			s.db[index] = v
			found = true
			break
		}
	}
	if !found {
		err := ErrRepositoryVehicleNotFound
		return err
	}
	err := s.loader.Write(s.db)
	if err != nil {
		return ErrRepositoryWritingStorage
	}
	return nil
}

func (s *VehicleSlice) UpdateVehicleFuelType(vehicleId int, newFuelType string) error {
	found := false
	for _, v := range s.db {
		if v.ID == vehicleId {
			index := slices.Index(s.db, v)
			v.Attributes.FuelType = newFuelType

			s.db[index] = v
			found = true
			break
		}
	}
	if !found {
		err := ErrRepositoryVehicleNotFound
		return err
	}
	err := s.loader.Write(s.db)
	if err != nil {
		return ErrRepositoryWritingStorage
	}
	return nil
}

func (s *VehicleSlice) DeleteVehicle(vehicleId int) error {
	found := false
	for _, v := range s.db {
		if v.ID == vehicleId {
			index := slices.Index(s.db, v)
			s.db = slices.Delete(s.db, index, index+1)
			found = true
			break
		}
	}
	if !found {
		err := ErrRepositoryVehicleNotFound
		return err
	}
	err := s.loader.Write(s.db)
	if err != nil {
		return ErrRepositoryWritingStorage
	}
	return nil
}
