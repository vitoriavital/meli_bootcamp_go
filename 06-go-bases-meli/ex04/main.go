package main

import (
	"errors"
	"fmt"
)

func main() {
	var Err error
	salary := 13
	LowSalaryErr := errors.New("Error: the minimum taxable amount is 150,000 and the salary entered is:")

	if salary < 150000 {
		Err = fmt.Errorf("%w %d", LowSalaryErr, salary)
	}

	if errors.Is(Err, LowSalaryErr) {
		fmt.Println(Err)
	} else {
		fmt.Println("Must pay tax")
	}
}