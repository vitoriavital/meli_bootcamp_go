package cmd

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id 			int
	Name 		string
	Email		string
	Dest		string
	FlightTime	string
	Price		float64
}

func loadDetails(lines []byte) ([]Ticket, error) {
	content := string(lines)
	allLines := strings.Split(content, "\n")
	var allTickets []Ticket
	
	for _, line := range allLines {
		words := strings.Split(line, ",")
		if len(words) == 6 {
			nId,err := strconv.Atoi(words[0])
			if err != nil {
				return nil, errors.New("Error: Can't convert ID to int")
			}
			nPrice, err := strconv.ParseFloat(words[5], 64)
			if err != nil {
				return nil, errors.New("Error: Can't convert price to float64")
			}
			customer := Ticket{
				Id: nId,
				Name: words[1],
				Email: words[2],
				Dest: words[3],
				FlightTime: words[4],
				Price: nPrice,
			}
			allTickets = append(allTickets, customer)
		}
	}
	return allTickets, nil
}

func loadTickets() ([]Ticket, error) {
	file, err := os.ReadFile("/Users/mlvital/Desktop/bootcamp/meli_bootcamp_go/desafio-go-bases/tickets.csv")
	if err != nil {
		return nil, errors.New("Error: Invalid File!")
	}
	return loadDetails(file)
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	allTickets, err := loadTickets()
	if err != nil {
		return 0, err
	}
	var amount int
	for _, ticket := range allTickets {
		if ticket.Dest == destination {
			amount++
		}
	}
	return amount, nil
}

// ejemplo 2
func GetCountByPeriod(time string) (int, error) {
	var hours []int
	if time == "early morning" {
		hours = append(hours, 0, 6)
	} else if time == "morning" {
		hours = append(hours, 7, 12)
	} else if time == "afternoon" {
		hours = append(hours, 13, 19)
	} else if time == "night" {
		hours = append(hours, 20, 23)
	} else {
		msg := fmt.Sprint("Error: ", time, " is not a valid period of time!")
		return 0, errors.New(msg)
	}
	allTickets, err := loadTickets()
	if err != nil {
		return 0, err
	}
	var amount int
	for _, ticket := range allTickets {
		flightHour := strings.Split(ticket.FlightTime, ":")
		hour, err := strconv.Atoi(flightHour[0])
		if err != nil {
			return 0, errors.New("Error: Can't convert flight hour to int")
		}
		if hour >= hours[0] && hour <= hours[1] {
			amount++
		}
	}
	return amount, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (float64, error) {
	totalTicketToDest, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	rate := (float64(totalTicketToDest) / float64(total)) * 100
	return math.Round(rate * 100) / 100, nil
}
