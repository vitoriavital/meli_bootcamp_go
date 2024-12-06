package main

import (
	"go-web/internal/handler"
	"go-web/internal/service"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
    rt := chi.NewRouter()

	myService := service.NewProductService("/Users/mlvital/Desktop/bootcamp/meli_bootcamp_go/11-go-web-post-method-meli/go-web/docs/db/products.json")
	myHandler := handler.NewProductHandler(myService)
	rt.Get("/ping", myHandler.HandlerPing)
	rt.Route("/products", func (r chi.Router) {
		r.Get("/", myHandler.HandlerProducts)
		r.Post("/", myHandler.HandlerCreateProduct)
		r.Get("/{id}", myHandler.HandlerProductById)
		r.Get("/search", myHandler.HandlerProductSearch)
	})

    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }  
}