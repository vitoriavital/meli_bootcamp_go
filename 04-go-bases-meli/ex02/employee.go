package main

import "fmt"

type Person struct {
	ID			int
	Name		string
	DateOfBirth	string
}

type Employee struct {
	ID			int
	Position	string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Employee's Name: %s\n", e.Name)
	fmt.Printf("Employee's Date of Birth: %s\n", e.DateOfBirth)
	fmt.Printf("Employee's Position: %s\n", e.Position)
}

func main() {
	person := Person{1, "John", "01/02/2002"}
	employee := Employee{1, "Tech Lead", person}

	employee.PrintEmployee()
}