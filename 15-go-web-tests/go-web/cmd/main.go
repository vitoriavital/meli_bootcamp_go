package main

import (
	"go-web/internal/products/controller"
	"go-web/internal/products/repository"
	"go-web/internal/products/service"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
)

func main() {
    rt := chi.NewRouter()
	err := godotenv.Load("API_TOKEN")
	if err != nil {
		log.Println("Failed to load token from .env file")
	}
	repo := repository.NewProductRepository("docs/db/products.json")
	productService := service.NewProductService(repo)
	productController := controller.NewProductController(productService)
	rt.Get("/ping", productController.HandlerPing)
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", productController.GetAllProducts)
		r.Post("/", productController.CreateProduct)
		r.Get("/{id}", productController.GetProductById)
		r.Put("/{id}", productController.UpdateProduct)
		r.Patch("/{id}", productController.PatchProduct)
		r.Delete("/{id}", productController.DeleteProduct)
		r.Get("/search", productController.SearchProduct)
	})
    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }  
}