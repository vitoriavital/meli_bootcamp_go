package main

import (
	"errors"
	"fmt"
)

type LowSalaryError struct {
	message string
}

func (e *LowSalaryError) Error() string {
	return e.message
}

func main() {
	var LowSalaryErr error = &LowSalaryError{"Error: salary is less than 10000"}
	var Err error
	salary := 13

	if salary <= 10000 {
		Err = fmt.Errorf("%w", LowSalaryErr)
	}

	if errors.Is(Err, LowSalaryErr) {
		fmt.Println("Err is LowSalaryErr")
	} else {
		fmt.Println("Err is not LowSalaryErr")
	}
}