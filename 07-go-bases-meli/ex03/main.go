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

func checkValues(user Data) (string, error) {
	if user.File == "" || user.Name == "" || user.ID == "" || user.PhoneNumber == "" || user.Adress == "" {
		return "", errors.New("Error: Empty field detected")
	}
	content := fmt.Sprintf("%s,%s,%s,%s,%s\n", user.File, user.Name, user.ID, user.PhoneNumber, user.Adress)

	return content, nil
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

	userContent, err := checkValues(newUser)
	if err != nil {
		return err
	}
	file,_ := os.OpenFile(newUser.File,  os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	_, err = file.Write([]byte(userContent))
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
	fmt.Scanln(&newUser.Name)
	fmt.Print("Type a ID: ")
	fmt.Scanln(&newUser.ID)
	fmt.Print("Type a phone number: ")
	fmt.Scanln(&newUser.PhoneNumber)
	fmt.Print("Type a adress: ")
	fmt.Scanln(&newUser.Adress)
}

func main() {
	newUser := Data{}
	defer endOfProgram()

	fmt.Print("Type a file: ")
	fmt.Scanln(&newUser.File)
	
	file, err :=os.ReadFile(newUser.File)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	
	askDetails(&newUser)
	currentCustomers := readDetails(file)
	err = saveDetails(newUser, currentCustomers)
	if err != nil {
		panic(err)
	}
	checkCreation(newUser.File)
}