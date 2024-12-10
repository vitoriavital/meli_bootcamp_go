package controller_test

import (
	"go-web/internal/products/controller"
	"go-web/internal/products/repository"
	"go-web/internal/products/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestGetAllProducts(t *testing.T) {
	t.Run("success to get all products", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()

		productController.GetAllProducts(res, req)
		expectedCode := http.StatusOK
		expectedBody := `
		{
			"message": "success to get all products",
			"products": [
				{
					"id": 346,
					"name": "Flour - Bran, Red",
					"quantity": 452,
					"code_value": "S93304S",
					"is_published": true,
					"expiration": "08/04/2021",
					"price": 990.64
				},
				{
					"id": 95,
					"name": "Sole - Dover, Whole, Fresh",
					"quantity": 90,
					"code_value": "S72392",
					"is_published": false,
					"expiration": "12/12/2021",
					"price": 196.64
				}
			],
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be OK")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
}


func TestGetProductById(t *testing.T) {
	t.Run("success to get product id 95", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products/95", nil)
		res := httptest.NewRecorder()

		productController.GetProductById(res, req)
		expectedCode := http.StatusOK
		expectedBody := `
		{
			"message": "product found",
			"product":
				{
					"id": 95,
					"name": "Sole - Dover, Whole, Fresh",
					"quantity": 90,
					"code_value": "S72392",
					"is_published": false,
					"expiration": "12/12/2021",
					"price": 196.64
				},
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be OK")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("invalid id format", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products/ok", nil)
		res := httptest.NewRecorder()

		productController.GetProductById(res, req)
		expectedCode := http.StatusBadRequest
		expectedBody := `
		{
			"message": "invalid value for id",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusBadRequest")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("non-existent id", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products/10", nil)
		res := httptest.NewRecorder()

		productController.GetProductById(res, req)
		expectedCode := http.StatusNotFound
		expectedBody := `
		{
			"message": "product not found",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusBadRequest")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestSearchProduct(t *testing.T) {
	t.Run("success to get products with price >= 990", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products/search?priceGt=990", nil)
		res := httptest.NewRecorder()

		productController.SearchProduct(res, req)
		expectedCode := http.StatusOK
		expectedBody := `
		{
			"message": "products found",
			"products": [
				{
					"id": 346,
					"name": "Flour - Bran, Red",
					"quantity": 452,
					"code_value": "S93304S",
					"is_published": true,
					"expiration": "08/04/2021",
					"price": 990.64
				}
			],
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be OK")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("missing price query param", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products/search", nil)
		res := httptest.NewRecorder()

		productController.SearchProduct(res, req)
		expectedCode := http.StatusBadRequest
		expectedBody := `
		{
			"message": "missing priceGt",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusBadRequest")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
	t.Run("no products in that price range", func(t *testing.T) {
		repo := repository.NewProductRepository("../../../docs/db/products_test.json")
		productService := service.NewProductService(repo)
		productController := controller.NewProductController(productService)

		req := httptest.NewRequest("GET", "/products/search?priceGt=1000", nil)
		res := httptest.NewRecorder()

		productController.SearchProduct(res, req)
		expectedCode := http.StatusInternalServerError
		expectedBody := `
		{
			"message": "no products found for this priceGt",
			"error": true
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code, "expected status code to be StatusInternalServerError")
		require.JSONEq(t, expectedBody, res.Body.String(), "response body mismatch")
		require.Equal(t, expectedHeader, res.Header())
	})
}
