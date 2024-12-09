package controller_test

import (
	"bytes"
	"go-web/internal/products/controller"
	"go-web/internal/products/repository"
	"go-web/internal/products/service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	err := godotenv.Load("API_TOKEN")
	if err != nil {
		log.Println("Failed to load token from .env file")
	}
	t.Run("success to create new product", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		r := chi.NewRouter()
		r.Post("/products", productController.CreateProduct)

		jsonProduct := `{
			"name": "chocolate",
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("POST", "/products", bytes.NewReader([]byte(jsonProduct)))

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusCreated
		expectedBody := `
		{
			"message": "new product created",
			"product":
				{
					"id": 3,
					"name": "chocolate",
					"quantity": 10,
					"code_value": "code943fg",
					"is_published": true,
					"expiration": "10/09/2024",
					"price": 9.5
				},
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusCreated")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to create new product because of missing field", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		r := chi.NewRouter()
		r.Post("/products", productController.CreateProduct)

		jsonProduct := `{
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("POST", "/products", bytes.NewReader([]byte(jsonProduct)))

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusBadRequest
		expectedBody := `
		{
			"message": "couldn't create product",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusBadRequest")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to create new product because of wrong token", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		os.Setenv("API_TOKEN", "123")
		r := chi.NewRouter()
		r.Post("/products", productController.CreateProduct)

		jsonProduct := `{
			"name": "chocolate",
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("POST", "/products", bytes.NewReader([]byte(jsonProduct)))

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusUnauthorized
		expectedBody := `
		{
			"message": "unauthorized",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusUnauthorized")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
}
