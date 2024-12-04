package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    // server
    rt := chi.NewRouter()
    // -> endpoints
    rt.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
        // set code and body
        w.WriteHeader(http.StatusOK) 
		w.Write([]byte("Hello World!"))
    })
    // run
    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }  
}
