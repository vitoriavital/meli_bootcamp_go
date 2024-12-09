package controller_test

import (
	"bytes"
	"go-web/internal/products/controller"
	"go-web/internal/products/repository"
	"go-web/internal/products/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"log"
)

func TestUpdateProduct(t *testing.T) {
	err := godotenv.Load("../../../.env")
    if err != nil {
        log.Println("Failed to load .env file")
        return
    }
	t.Run("success to update product", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Put("/products/{id}", productController.UpdateProduct)

		jsonProduct := `{
			"name": "strawberry",
			"quantity": 10,
			"code_value": "123",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("PUT", "/products/3", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusOK
		expectedBody := `
		{
			"message": "product updated",
			"product":
				{
					"id": 3,
					"name": "strawberry", 
					"quantity": 10,
					"code_value": "123",
					"is_published": true,
					"expiration": "10/09/2024",
					"price": 9.5
				},
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusOK")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to update product because of invalid id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		r := chi.NewRouter()
		r.Put("/products/{id}", productController.UpdateProduct)

		jsonProduct := `{
			"name": "strawberry", 
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("PUT", "/products/abc", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusBadRequest
		expectedBody := `
		{
			"message": "impossible conversion of id to int",
			"error": true
		}`
		t.Logf("Response Body: %s", res.Body.String())
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusBadRequest")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to update product because of non-existent id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Put("/products/{id}", productController.UpdateProduct)

		jsonProduct := `{
			"name": "chocolate", 
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("PUT", "/products/78", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusNotFound
		expectedBody := `
		{
			"message": "product not found",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusNotFound")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to update product because of wrong token", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Put("/products/{id}", productController.UpdateProduct)
		jsonProduct := `{
			"name": "chocolate", 
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("PUT", "/products/346", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "wrong-key")
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


func TestPatchProduct(t *testing.T) {
	err := godotenv.Load("../../../.env")
    if err != nil {
        log.Println("Failed to load .env file")
        return
    }
	t.Run("success to patch product", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		r := chi.NewRouter()
		r.Patch("/products/{id}", productController.PatchProduct)

		jsonProduct := `{
			"code_value": "some_code"
		}`

		req := httptest.NewRequest("PATCH", "/products/3", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusOK
		expectedBody := `
		{
			"message": "product updated",
			"product":
				{
					"id": 3,
					"name": "strawberry", 
					"quantity": 10,
					"code_value": "some_code",
					"is_published": true,
					"expiration": "10/09/2024",
					"price": 9.5
				},
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusOK")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to patch product because of invalid id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		r := chi.NewRouter()
		r.Patch("/products/{id}", productController.PatchProduct)

		jsonProduct := `{
			"code_value": "S93304S"
		}`

		req := httptest.NewRequest("PATCH", "/products/abc", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusBadRequest
		expectedBody := `
		{
			"message": "impossible conversion of id to int",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusBadRequest")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to patch product because of non-existent id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Patch("/products/{id}", productController.UpdateProduct)

		jsonProduct := `{
			"name": "chocolate", 
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("PATCH", "/products/78", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusNotFound
		expectedBody := `
		{
			"message": "product not found",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusNotFound")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to patch product because of wrong token", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Patch("/products/{id}", productController.UpdateProduct)
		jsonProduct := `{
			"name": "chocolate", 
			"quantity": 10,
			"code_value": "code943fg",
			"is_published": true,
			"expiration": "10/09/2024",
			"price": 9.5
		}`

		req := httptest.NewRequest("PATCH", "/products/346", bytes.NewReader([]byte(jsonProduct)))
		req.Header.Set("API_TOKEN", "wrong-key")
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