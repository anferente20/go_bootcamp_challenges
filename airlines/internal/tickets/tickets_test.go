package tickets_test

import (
	"acmevision-airlines/internal/constants"
	"acmevision-airlines/internal/tickets"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTicketsByDestination(t *testing.T) {

	t.Run("Validate get Ticket by destination SUCCESS",
		func(t *testing.T) {
			tickets.GetTickets(constants.TicketsFile)

			result := 18
			const country = "Colombia"

			testResult, _ := tickets.GetTotalTicketsByDestination(country)
			assert.Equal(t, result, testResult, "Result should be 18")

		})

	t.Run("Validate get Ticket by destination FAILURE",
		func(t *testing.T) {
			result := "Destination not registered!"
			const country = "Pikachu"

			_, errorResult := tickets.GetTotalTicketsByDestination(country)
			assert.Equal(t, result, errorResult.Error(), "Result should be 18")

		})
}

func TestGetCountByPeriod(t *testing.T) {

	t.Run("Validate get Tickets by Early Morning SUCCESS",
		func(t *testing.T) {
			tickets.GetTickets(constants.TicketsFile)

			result := 304

			testResult, _ := tickets.GetCountByPeriod(constants.EarlyMorning)
			assert.Equal(t, result, testResult, "Result should be 304")

		})

	t.Run("Validate get Tickets by Morning SUCCESS",
		func(t *testing.T) {

			result := 256

			testResult, _ := tickets.GetCountByPeriod(constants.Morning)
			assert.Equal(t, result, testResult, "Result should be 256")

		})

	t.Run("Validate get Tickets by Afternoon SUCCESS",
		func(t *testing.T) {

			result := 289

			testResult, _ := tickets.GetCountByPeriod(constants.Afternoon)
			assert.Equal(t, result, testResult, "Result should be 289")

		})

	t.Run("Validate get Tickets by Night SUCCESS",
		func(t *testing.T) {

			result := 151

			testResult, _ := tickets.GetCountByPeriod(constants.Night)
			assert.Equal(t, result, testResult, "Result should be 151")

		})
}

func TestAverageDestination(t *testing.T) {

	t.Run("Validate get Ticket by destination SUCCESS",
		func(t *testing.T) {
			tickets.GetTickets(constants.TicketsFile)

			result := float32(0.018)
			const country = "Colombia"

			testResult, _ := tickets.AverageDestination(country)
			assert.Equal(t, result, testResult, "Result should be 0.018")

		})

	t.Run("Validate get Ticket by destination FAILURE",
		func(t *testing.T) {
			result := "Destination not registered!"
			const country = "Pikachu"

			_, errorResult := tickets.AverageDestination(country)
			assert.Equal(t, result, errorResult.Error(), "Result should be 18")

		})
}
