package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	File		string
	Name		string
	ID			string
	PhoneNumber	string
	Adress		string
}

func checkCreation(file string) {
	fmt.Println("List after creation of new user: ")
	content, _ :=os.ReadFile(file)
	lines := string(content)
	fmt.Println(lines)
}

func readDetails(lines []byte) []Data {
	content := string(lines)
	allLines := strings.Split(content, "\n")
	var currentCustomers []Data
	
	for _, line := range allLines {
		words := strings.Split(line, ",")
		if len(words) == 5 {
			customer := Data{
				File: words[0],
				Name: words[1],
				ID: words[2],
				PhoneNumber: words[3],
				Adress: words[4],
			}
			currentCustomers = append(currentCustomers, customer)
		}
	}
	return currentCustomers
}

func saveDetails(newUser Data, allUsers []Data) error {
	for _, user := range allUsers {
		if user == newUser {
			return errors.New("Error: client already exists")
		}
	}
	addUser := fmt.Sprintf("%s,%s,%s,%s,%s\n", newUser.File, newUser.Name, newUser.ID, newUser.PhoneNumber, newUser.Adress)
	var err error
	file,_ := os.OpenFile(newUser.File,  os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	_, err = file.Write([]byte(addUser))
	if err != nil {
		return err
	}
	return nil
}

func endOfProgram() {
	err := recover()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Several errors were detected at runtime")
	}
	fmt.Println("End of execution")
}

func askDetails(newUser *Data) {
	fmt.Print("Type a name: ")
	fmt.Scan(&newUser.Name)
	fmt.Print("Type a ID: ")
	fmt.Scan(&newUser.ID)
	fmt.Print("Type a phone number: ")
	fmt.Scan(&newUser.PhoneNumber)
	fmt.Print("Type a adress: ")
	fmt.Scan(&newUser.Adress)
}

func main() {
	newUser := Data{}
	defer endOfProgram()

	fmt.Print("Type a file: ")
	fmt.Scan(&newUser.File)
	
	file, err :=os.ReadFile(newUser.File)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	
	askDetails(&newUser)
	currentCustomers := readDetails(file)
	if saveDetails(newUser, currentCustomers) != nil {
		panic("Error: client already exists")
	}
	checkCreation(newUser.File)
}