package main

import "fmt"

func main () {
	var word string
	fmt.Print("Type a word: ")
	fmt.Scan(&word)
	for _, letter := range word {
		fmt.Printf("%c\n", letter)
	}
}