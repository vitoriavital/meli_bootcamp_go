package controller_test

import (
	"go-web/internal/products/controller"
	"go-web/internal/products/repository"
	"go-web/internal/products/middleware"
	"go-web/internal/products/service"
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/go-chi/chi/v5"
	"os"
)

func TestDeleteProduct(t *testing.T) {
	t.Run("success to delete product", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		r := chi.NewRouter()
		r.Use(middleware.Auth)
		r.Delete("/products/{id}", productController.DeleteProduct)

		req := httptest.NewRequest("DELETE", "/products/3", nil)
		req.Header.Set("API_TOKEN",  "super-secure-key")
		os.Setenv("API_TOKEN", "super-secure-key")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		expectedCode := http.StatusOK
		expectedBody := `
		{
			"message": "product deleted",
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusOK")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("failed to delete product because of invalid id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Use(middleware.Auth)
		r.Delete("/products/{id}", productController.DeleteProduct)
		req := httptest.NewRequest("DELETE", "/products/ok", nil)
		req.Header.Set("API_TOKEN",  "super-secure-key")
		os.Setenv("API_TOKEN", "super-secure-key")
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
	t.Run("failed to delete product because of non-existent id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)
		r := chi.NewRouter()
		r.Use(middleware.Auth)
		r.Delete("/products/{id}", productController.DeleteProduct)
		req := httptest.NewRequest("DELETE", "/products/10", nil)
		req.Header.Set("API_TOKEN",  "super-secure-key")
		os.Setenv("API_TOKEN", "super-secure-key")
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
}
