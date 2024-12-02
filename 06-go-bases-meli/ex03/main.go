package main

import (
	"errors"
	"fmt"
)

func main() {
	var Err error
	salary := 13
	LowSalaryErr := errors.New("Error: salary is less than 10000")

	if salary <= 10000 {
		Err = fmt.Errorf("%w", LowSalaryErr)
	}

	if errors.Is(Err, LowSalaryErr) {
		fmt.Println("Err is LowSalaryErr")
	} else {
		fmt.Println("Err is not LowSalaryErr")
	}
}