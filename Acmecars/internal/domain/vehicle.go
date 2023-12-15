package domain

// VehicleAttributes is an struct that represents the attributes of a vehicle.
type VehicleAttributes struct {
	// Brand is the brand of the vehicle.
	Brand string `json:"brand"`
	// Model is the model of the vehicle.
	Model string `json:"model"`
	// Registration is the registration of the vehicle.
	Registration string `json:"registration"`
	// Year is the fabrication year of the vehicle.
	Year int `json:"year"`
	// Color is the color of the vehicle.
	Color string `json:"color"`

	// MaxSpeed is the maximum speed of the vehicle.
	MaxSpeed int `json:"max_speed"`
	// FuelType is the fuel type of the vehicle.
	FuelType string `json:"fuel_type"`
	// Transmission is the transmission of the vehicle.
	Transmission string `json:"transmission"`

	// Passengers is the capacity of passengers of the vehicle.
	Passengers int `json:"passengers"`

	// Height is the height of the vehicle.
	Height float64 `json:"height"`
	// Width is the width of the vehicle.
	Width float64 `json:"width"`

	// Weight is the weight of the vehicle.
	Weight float64 `json:"weight"`
}

// Vehicle is an struct that represents a vehicle.
type Vehicle struct {
	// ID is the unique identifier of the vehicle.
	ID int `json:"id"`
	// Attributes is the attributes of the vehicle.
	Attributes VehicleAttributes `json:"attributes"`
}
