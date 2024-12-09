package controller_test

import (
	"testing"
)

func TestOrderTestSuite(t *testing.T) {
	// go test create_test.go
    t.Run("TestGetAllProducts", TestGetAllProducts)
    t.Run("TestGetProductById", TestGetProductById)
    t.Run("TestCreateProduct", TestCreateProduct)
    t.Run("TestDeleteProduct", TestDeleteProduct)
}