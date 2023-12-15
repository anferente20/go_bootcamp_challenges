package application

import (
	"app/cmd/application/handler"
	"app/cmd/application/middlewares"
	"app/internal/vehicles/repository"
	"app/internal/vehicles/service"
	"app/pkg/loader"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// ConfigDefaultInMemory is an struct that contains the configuration for the default application settings.
type ConfigDefaultInMemory struct {
	// FileLoader is the path to the file that contains the vehicles.
	FileLoader string
	// Addr is the address where the application will be listening.
	Addr string
}

// NewDefaultInMemory returns a new instance of a default application.
func NewDefaultInMemory(c *ConfigDefaultInMemory) *DefaultInMemory {
	// default config
	defaultCfg := &ConfigDefaultInMemory{
		FileLoader: "vehicles.json",
		Addr:       ":8080",
	}
	if c != nil {
		if c.FileLoader != "" {
			defaultCfg.FileLoader = c.FileLoader
			fmt.Println(defaultCfg.FileLoader)
		}
		if c.Addr != "" {
			defaultCfg.Addr = c.Addr
		}
	}

	return &DefaultInMemory{
		fileLoader: defaultCfg.FileLoader,
		addr:       defaultCfg.Addr,
	}
}

// DefaultInMemory is an struct that contains the default application settings.
type DefaultInMemory struct {
	// fileLoader is the path to the file that contains the vehicles.
	fileLoader string
	// addr is the address where the application will be listening.
	addr string
}

// Run starts the application.
func (d *DefaultInMemory) Run() (err error) {
	// authenticator
	au := middlewares.NewAuthenticator(os.Getenv("TOKEN"))
	// dependencies initialization
	// loader
	ld := loader.NewVehicleJSON(d.fileLoader)

	// repository
	rp := repository.NewVehicleSlice(*ld)

	// service
	sv := service.NewDefault(rp)

	// handler
	hd := handler.NewVehicleDefault(sv)

	// router
	rt := gin.New()
	// - middlewares
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery())
	// - endpoints
	gr := rt.Group("/vehicles")
	{

		gr.GET("", hd.GetAll())
		gr.GET("/color/:color/year/:year", hd.SearchByColorAndYear())
		gr.GET("/brand/:brand/between/:start_year/:end_year", hd.SearchByBrandAndYear())
		gr.GET("/weight", hd.SearchVehicleByWeight())
		gr.GET("/dimensions", hd.SearchVehiclesByDimensions())
		gr.GET("/average_speed/brand/:brand", hd.AverageSpeedByBrand())
		gr.GET("/average_capacity/brand/:brand", hd.AverageCapacityByBrand())
		gr.GET("/fuel_type/:type", hd.SearchVehicleByFuelType())
		gr.GET("/transmission/:type", hd.SearchVehicleByTransmissionType())

		gr.POST("", hd.AddVehicle())
		gr.POST("/batch", hd.AddVehicles())

		gr.PUT("/:id/update_speed", hd.UpdateVehicleSpeed())
		gr.PUT("/:id/update_fuel", hd.UpdateVehicleFuelType())

		gr.DELETE("/:id", hd.DeleteVehicle())
	}

	gr.Use(au.Authenticate())
	// run application
	err = rt.Run(d.addr)
	if err != nil {
		return
	}

	return
}
