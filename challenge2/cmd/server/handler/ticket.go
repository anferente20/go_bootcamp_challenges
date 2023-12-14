package handler

import (
	"challenge2/internal/tickets"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service tickets.Service
}

func NewHandler(s tickets.Service) *TicketHandler {
	return &TicketHandler{
		service: s,
	}
}

// Show GetTicketsByCountry godoc
// @Summary Show All tickets by country
// @Description Getl all tickets by country
// @Tags ticjets
// @Produces json
// @Failure 500 {object} error
// @Success 200 {object}  map[string]string
// @Router /tickets/getByCountry/:dest [get]
func (s *TicketHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

// Show AverageDestination godoc
// @Summary Show average of tickets by country
// @Description Getl all tickets by country
// @Tags ticjets
// @Produces json
// @Failure 500 {object}  error
// @Success 200 {object}  map[string]string
// @Router /tickets/getAverage/:dest [get]
func (s *TicketHandler) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
