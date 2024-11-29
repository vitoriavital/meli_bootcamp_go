package main

import "fmt"

type Student struct {
	Name	string
	Surname	string
	ID		int
	Date	string
}

func (s Student) details() {
	fmt.Printf("%2s | %10s | %10s | %10s\n", "ID", "Name", "Surname", "Date")
	fmt.Printf("%2d | %10s | %10s | %8s\n", s.ID, s.Name, s.Surname, s.Date)
}

func main() {
	student := Student{"Vitória", "Vital", 1, "29/11/2024"}
	student.details()
}