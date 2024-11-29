package main

import "fmt"

func calculateSalary(minutes int, category string) float64 {
	var salary float64

	if category == "A" {
		salary = 3000.0 * (float64(minutes) / 60.0)
		return salary + 0.5 * salary
	} else if category == "B" {
		salary = 1500.0 * (float64(minutes) / 60.0)
		return salary + 0.2 * salary
	} else {
		return 1000.0 * (float64(minutes) / 60.0)
	}
}

func main() {
	var minutes		int
	var	category	string

	fmt.Print("Type amount of minutes you've worked this month: ")
	fmt.Scan(&minutes)
	fmt.Print("Type your category (A, B, or C): ")
	fmt.Scan(&category)

	fmt.Println("Total salary: US$ ", calculateSalary(minutes, category))
}