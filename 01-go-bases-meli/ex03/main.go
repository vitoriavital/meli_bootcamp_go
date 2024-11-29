package main

import "fmt"

func main () {
	var monthNumber int
	fmt.Print("Type the number of a month: ")
	fmt.Scan(&monthNumber)
	if monthNumber < 1 && monthNumber > 12 { 
		fmt.Println("Invalid number!")
		return
	} 

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Printf("%d, %s\n", monthNumber, months[monthNumber - 1])
}