package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
	"go-web/internal/model"
	"go-web/internal/service"
	"github.com/go-chi/chi/v5"
	"strconv"
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


func (h *ProductHandler) HandlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var requestBody model.RequestBodyProduct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't Create product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	newProduct := model.Product{
		Name:			requestBody.Name,
		Quantity:		requestBody.Quantity,
		CodeValue:		requestBody.CodeValue,
		IsPublished:	requestBody.IsPublished,
		Expiration:		requestBody.Expiration,
		Price:			requestBody.Price,
	}

	res, err := h.Service.CreateProduct(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't Create product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) HandlerProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	products, err := h.Service.GetAllProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	var msg string
	for _, p := range products {
		msg += fmt.Sprint("id: ", p.Id, " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}

func (h *ProductHandler) HandlerProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	var msg string
	msg += fmt.Sprint("id ", id, " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	w.Write([]byte(msg))
}

func (h *ProductHandler) HandlerProductSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	price := r.URL.Query().Get("priceGt")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Missing priceGt"))
		return
	}
	priceGt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	products,err := h.Service.GetProductsByPrice(priceGt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	var msg string
	for _, p := range products {
		msg += fmt.Sprint("id: ", p.Id , " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}