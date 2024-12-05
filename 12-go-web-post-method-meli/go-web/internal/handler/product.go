package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
	"go-web/internal/model"
	"go-web/internal/repository"
	"go-web/internal/service"
	"go-web/pkg/validations"
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
	validProduct := validations.ValidateNewProduct(requestBody)
	products, pErr := repository.LoadProducts()
	if err != nil || validProduct != nil || pErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			fmt.Println(err)
		}
		if validProduct != nil {
			fmt.Println(validProduct)
		}
		if pErr != nil {
			fmt.Println(pErr)
		}
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't Create product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	newProduct := model.Product{
		Id: 			len(products) + 1,
		Name:			requestBody.Name,
		Quantity:		requestBody.Quantity,
		CodeValue:		requestBody.CodeValue,
		IsPublished:	requestBody.IsPublished,
		Expiration:		requestBody.Expiration,
		Price:			requestBody.Price,
	}
	res := model.ResponseBodyProduct{
		Message: "New product created!",
		Product: &newProduct,
		Error: false,
	}
	err = repository.SaveProducts(append(products, newProduct))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't save product!",
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
	products, err := repository.LoadProducts()
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
	products,err := repository.LoadProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := products[nId - 1]
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
	products,err := repository.LoadProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	var allProducts []model.Product
	for _, product := range products {
		if product.Price >= priceGt {
			allProducts = append(allProducts, product)
		}
	}
	var msg string
	for _, p := range allProducts {
		msg += fmt.Sprint("id: ", p.Id , " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}