package main

import (
	"errors"
	"fmt"
)

func calculateSalary(hours int, rate int) (int, error) {
	if hours < 80 || hours < 0 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	salary := hours * rate
	if salary < 150000 {
		LowSalaryErr := errors.New("Error: the minimum taxable amount is 150,000 and the salary entered is:")
		Err := fmt.Errorf("%w %d", LowSalaryErr, salary)
		return 0, Err
	}
	return salary, nil
}

func main() {
	var hours int
	var rate int

	fmt.Print("Type amount of hours you've worked this month: ")
	fmt.Scan(&hours)
	fmt.Print("Type your hourly rate: ")
	fmt.Scan(&rate)

	salary, err := calculateSalary(hours, rate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Must pay taxes")
		fmt.Printf("Salary before taxes: $%d | Salary after taxes(10%%) $%.02f", salary, 0.9 * float64(salary))
	}
}