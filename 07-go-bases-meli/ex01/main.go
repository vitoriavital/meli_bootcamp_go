package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("execução concluída")
	} ()
	_, err :=os.ReadFile("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
}