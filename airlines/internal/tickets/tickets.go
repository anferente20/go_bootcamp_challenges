package tickets

import (
	"acmevision-airlines/internal/constants"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id          int
	name        string
	email       string
	destination string
	time        string
	price       float32
}

var Countrys map[string]int
var Tickets []Ticket

func GetTickets(fileName string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	Countrys = make(map[string]int)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6

	records, err := reader.ReadAll()
	if err != nil {
		panic("Error reading records")
	}

	for _, record := range records {
		ticketId, er := strconv.Atoi(record[0])
		if er != nil {
			fmt.Println(er)
		}
		var price int
		price, er = strconv.Atoi(record[5])
		if er != nil {
			fmt.Println(er)
		}
		Tickets = append(Tickets, Ticket{id: ticketId, name: record[1], email: record[2], destination: record[3], time: record[4], price: float32(price)})

		if ticketCount, hasTickets := Countrys[record[3]]; hasTickets {
			Countrys[record[3]] = ticketCount + 1
		} else {
			Countrys[record[3]] = 1
		}
	}
}

// ejemplo 1
func GetTotalTicketsByDestination(destination string) (int, error) {
	if ticketCount, hasTickets := Countrys[destination]; hasTickets {
		return ticketCount, nil
	}
	return 0, errors.New("Destination not registered!")
}

// // ejemplo 2
func GetCountByPeriod(time string) (int, error) {
	countByPeriod := 0
	for _, ticket := range Tickets {
		if belongsToPeriod(ticket, time) {
			countByPeriod++
		}
	}
	if countByPeriod == 0 {
		return 0, errors.New("Period has no flights")
	}
	return countByPeriod, nil
}

// // ejemplo 3
func AverageDestination(destination string) (float32, error) {
	if ticketCount, hasTickets := Countrys[destination]; hasTickets {

		return float32(ticketCount) / float32(len(Tickets)), nil
	}
	return 0, errors.New("Destination not registered!")
}

func belongsToPeriod(ticket Ticket, time string) bool {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	switch time {
	case constants.EarlyMorning:
		hour, err := strconv.Atoi(strings.Split(ticket.time, ":")[0])
		if err != nil {
			panic("Wrong hour")
		}
		return validateTimeFrame(hour, 0, 6)
	case constants.Morning:
		hour, err := strconv.Atoi(strings.Split(ticket.time, ":")[0])
		if err != nil {
			panic("Wrong hour")
		}
		return validateTimeFrame(hour, 7, 12)
	case constants.Afternoon:
		hour, err := strconv.Atoi(strings.Split(ticket.time, ":")[0])
		if err != nil {
			panic("Wrong hour")
		}
		return validateTimeFrame(hour, 13, 19)
	case constants.Night:
		hour, err := strconv.Atoi(strings.Split(ticket.time, ":")[0])
		if err != nil {
			panic("Wrong hour")
		}
		return validateTimeFrame(hour, 20, 23)
	default:
		hour, err := strconv.Atoi(strings.Split(ticket.time, ":")[0])
		if err != nil {
			panic("Wrong hour")
		}
		return validateTimeFrame(hour, 0, 6)
	}
}

func validateTimeFrame(hour, startHourFrme, lastHourFrame int) bool {
	if hour >= startHourFrme && hour <= lastHourFrame {
		return true
	} else {
		return false
	}
}
