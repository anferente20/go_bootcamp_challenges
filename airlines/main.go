package main

import (
	"acmevision-airlines/internal/constants"
	"acmevision-airlines/internal/tickets"
	"bufio"
	"fmt"
	"os"
)

func main() {
	tickets.GetTickets("tickets.csv")

	fmt.Print("Wich country do you want to consult? ")
	scanner := bufio.NewScanner(os.Stdin)
	country := ""
	if scanner.Scan() {
		country = scanner.Text()
	}

	total, err := tickets.GetTotalTicketsByDestination(country)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Tickets to %s: %d \n", country, total)

	countEarly, err := tickets.GetCountByPeriod(constants.EarlyMorning)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Early Morning Flights: ", countEarly)

	countMorning, err := tickets.GetCountByPeriod(constants.Morning)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Morning Flights: ", countMorning)

	afternoon, err := tickets.GetCountByPeriod(constants.Afternoon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Afternoon Flights: ", afternoon)
	countNight, err := tickets.GetCountByPeriod(constants.Night)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Night Flights: ", countNight)

	average, err := tickets.AverageDestination("Colombia")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Percentage Tickets to Colombia: ", average, "%")

}
