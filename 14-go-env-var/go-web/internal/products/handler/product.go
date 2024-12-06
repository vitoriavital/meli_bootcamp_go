package handler

import (
	"net/http"
	"go-web/internal/products/service"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) HandlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}