package main

import "fmt"

type SalaryError struct {
	message string
}

func (e *SalaryError) Error() string {
	return e.message
}

func main() {
	salary := 160000
	
	if salary < 150000 {
		var Err error = &SalaryError{"Error: the salary entered does not reach the taxable minimum"}
		fmt.Println(Err)
		return
	}

	fmt.Println("Must pay tax")
}