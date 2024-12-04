package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName	string	`json:"first-name"`
	LastName	string	`json:"last-name"`
}

func main() {
	p := Person{
		FirstName: "Maria",
		LastName: "Vital",
	}
    rt := chi.NewRouter()
    rt.Get("/greetings", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) 
		msg := fmt.Sprint("Hello ", p.FirstName, " ",  p.LastName)
		w.Write([]byte(msg))
    })
    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }  
}