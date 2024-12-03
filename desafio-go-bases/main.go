package main

import (
	"fmt"
	t "github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	total, err := t.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total tickets to Brazil:", total)

	morning, err := t.GetCountByPeriod("early morning")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total morning tickets:", morning)

	percentage, err := t.AverageDestination("France", 1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Percentage of tickets to France: %.02f%%", percentage)
}
