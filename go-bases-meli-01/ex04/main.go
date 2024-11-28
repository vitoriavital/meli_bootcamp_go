package main

import "fmt"

func main () {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Benjamin's age: %d\n", employees["Benjamin"])

	var legalAgeEmployees int
	for _, employeeAge := range employees {
		if employeeAge > 21 {
			legalAgeEmployees++
		}
	}
	fmt.Printf("Employees with more than 21 years: %d\n", legalAgeEmployees)

	fmt.Println("Creating Federico")
	employees["Federico"] = 25
	if _, exists := employees["Pedro"]; exists {
		fmt.Println("Deleting Pedro")
		delete(employees, "Pedro")
	}
}