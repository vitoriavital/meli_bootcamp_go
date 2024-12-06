package main

import (
	"go-web/internal/handler"
	"go-web/internal/service"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
    rt := chi.NewRouter()

	productService := service.NewProductService("products.json")
	myHandler := handler.NewProductHandler(productService)
	rt.Get("/ping", myHandler.HandlerPing)
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", myHandler.HandlerProducts)
		r.Post("/", myHandler.HandlerCreateProduct)
		r.Get("/{id}", myHandler.HandlerProductById)
		r.Put("/{id}", myHandler.HandlerUpdateProduct)
		r.Patch("/{id}", myHandler.HandlerPatchProduct)
		r.Delete("/{id}", myHandler.HandlerDeleteProduct)
		r.Get("/search", myHandler.HandlerProductSearch)
	})
    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }  
}