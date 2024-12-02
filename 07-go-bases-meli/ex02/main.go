package main

import (
	"fmt"
	"os"
)

func readDetails(lines []byte) {
	content := string(lines)
	fmt.Println(content)
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("execução concluída")
	} ()
	file, err :=os.ReadFile("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	readDetails(file)
}