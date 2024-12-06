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
		msg += fmt.Sprint("id: ", p.Id, " - ", p.Name, " - $", p.Price, " - ", p.Quantity, "\n")
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
	if err != nil || p == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Write([]byte("Error: This product doesn't exist!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	var msg string
	msg += fmt.Sprint("id ", id, " - ", p.Name, " - $", p.Price, " - ", p.Quantity, "\n")
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
		msg += fmt.Sprint("id: ", p.Id , " - ", p.Name, " - $", p.Price, " - ", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}

func (h *ProductHandler) HandlerUpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	err.Error(),
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	err.Error(),
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
	}
	var requestBody model.RequestBodyProduct
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't update or create product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	product := model.Product{
		Name:			requestBody.Name,
		Quantity:		requestBody.Quantity,
		CodeValue:		requestBody.CodeValue,
		IsPublished:	requestBody.IsPublished,
		Expiration:		requestBody.Expiration,
		Price:			requestBody.Price,
	}
	var res *model.ResponseBodyProduct
	if p != nil {
		res, err = h.Service.UpdateProduct(product, p.Id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			responseBody := model.ResponseBodyProduct{
				Message:	"Couldn't update existing product!",
				Product: nil,
				Error: true,
			}
			json.NewEncoder(w).Encode(responseBody)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
		return
	}
	res, err = h.Service.CreateProduct(product)
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

func patchValueHelperStr(reqValue *string, value string) string {
	if reqValue != nil {
		return *reqValue
	} else {
		return value
	}
}

func patchValueHelperInt(reqValue *int, value int) int {
	if reqValue != nil {
		return *reqValue
	} else {
		return value
	}
}

func patchValueHelperFloat(reqValue *float64, value float64) float64 {
	if reqValue != nil {
		return *reqValue
	} else {
		return value
	}
}

func patchValueHelperBool(reqValue *bool, value bool) bool {
	if reqValue != nil {
		return *reqValue
	} else {
		return value
	}
}

func (h *ProductHandler) HandlerPatchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	err.Error(),
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
	}
	if p != nil {
		fmt.Println("Error: Product not found!")
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	"Error: Product not found!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
	}
	var requestBody model.RequestUpdateBodyProduct
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't update product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	var product model.Product
	product.Id = p.Id
	product.Name = patchValueHelperStr(requestBody.Name, p.Name)
	product.Quantity = patchValueHelperInt(requestBody.Quantity, p.Quantity)
	product.CodeValue = patchValueHelperStr(requestBody.CodeValue, p.CodeValue)
	product.IsPublished = patchValueHelperBool(requestBody.IsPublished, p.IsPublished)
	product.Expiration = patchValueHelperStr(requestBody.Expiration, p.Expiration)
	product.Price = patchValueHelperFloat(requestBody.Price, p.Price)
	var res *model.ResponseBodyProduct
	res, err = h.Service.UpdateProduct(product, p.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't update existing product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}

func (h *ProductHandler) HandlerDeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	err.Error(),
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	err.Error(),
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	if p == nil {
		fmt.Println("Error: Product not found!")
		w.WriteHeader(http.StatusBadRequest)
		responseBody := model.ResponseBodyProduct{
			Message:	"Error: Product not found!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	var res *model.ResponseBodyProduct
	res, err = h.Service.DeleteProduct(p.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := model.ResponseBodyProduct{
			Message:	"Couldn't delete existing product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}