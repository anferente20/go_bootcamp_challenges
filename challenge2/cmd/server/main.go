// swagger: "2.0"
// info:
//
//	title: Tickets API
//	description: Spec Documentation for pet service.
//	version: 1.0.0
//
// schemes:
//   - http
//
// BasePath: /
// Host: localhost:8080
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"challenge2/cmd/server/middleware"
	"challenge2/cmd/server/router"
	"challenge2/internal/domain"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger documentation
// @title Tickets API
// @description this is a rest api server for items
// @version 1
// @address localhost:8080

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}
	lg := middleware.NewLogger()
	//define router
	r := gin.New()
	// > middlewares
	r.Use(lg.Log())
	r.Use(gin.Recovery())

	// >> docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	router := router.NewRouter(r, list)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
