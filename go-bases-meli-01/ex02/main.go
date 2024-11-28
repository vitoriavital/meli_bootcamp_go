package main

import "fmt"

func main () {
	var age int
	var isEmployed bool = false
	var yearsEmployed int
	var salary int

	fmt.Print("What is your age? ")
	fmt.Scan(&age)
	fmt.Print("Are you employed? Choose yes = 1 or no = 0: ")
	fmt.Scan(&isEmployed)
	if isEmployed == true {
		fmt.Print("How long have you been employed? ")
		fmt.Scan(&yearsEmployed)
		fmt.Print("How much is your salary: ")
		fmt.Scan(&salary)
	}
	
	if age > 22 && isEmployed && yearsEmployed > 1 {
		if salary > 100000 {
			fmt.Println("Loan denied! You already have a lot of money!")
		} else {
			fmt.Println("Loan approved! Go get your money!")
		}
	} else {
		fmt.Println("Sorry, you don't meet the requirements to get this loan :(")
	}
}