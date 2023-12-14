package router

import (
	"challenge2/cmd/server/handler"
	"challenge2/cmd/server/middleware"
	"challenge2/internal/domain"
	"challenge2/internal/tickets"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	list    []domain.Ticket
	handler handler.TicketHandler
}

func NewRouter(router *gin.Engine, list []domain.Ticket) Router {

	//add dependencies
	repo := tickets.NewRepository(list)
	service := tickets.NewService(repo)
	ticketHandler := handler.NewHandler(service)

	return Router{router: router, list: list, handler: *ticketHandler}
}

func (r *Router) MapRoutes() {
	// get Token
	token := os.Getenv("TOKEN")
	au := middleware.NewAuthenticator(token)

	grTickets := r.router.Group("/ticket")
	grTickets.Use(au.Authenticate())

	grTickets.GET("/getByCountry/:dest", r.handler.GetTicketsByCountry())
	grTickets.GET("/getAverage/:dest", r.handler.AverageDestination())
}
