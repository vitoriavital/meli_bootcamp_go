package main

import (
	"go-web/internal/products/controller"
	"go-web/internal/products/repository"
	"go-web/internal/products/service"
	"go-web/internal/products/middleware"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
)

func main() {
    rt := chi.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load token from .env file")
		return
	}
	repo := repository.NewProductRepository("docs/db/products.json")
	productService := service.NewProductService(repo)
	productController := controller.NewProductController(productService)
	rt.Use(middleware.MyLogger)
	rt.Get("/ping", productController.HandlerPing)
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", productController.GetAllProducts)
		r.Get("/{id}", productController.GetProductById)
		r.Get("/search", productController.SearchProduct)
		r.With(middleware.Auth).Post("/", productController.CreateProduct)
		r.With(middleware.Auth).Put("/{id}", productController.UpdateProduct)
		r.With(middleware.Auth).Patch("/{id}", productController.PatchProduct)
		r.With(middleware.Auth).Delete("/{id}", productController.DeleteProduct)
	})
    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }
}
