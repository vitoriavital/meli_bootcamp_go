package controller

import (
	"net/http"
	"go-web/internal/products/service"
)

type ProductController struct {
	Service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (h *ProductController) HandlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
