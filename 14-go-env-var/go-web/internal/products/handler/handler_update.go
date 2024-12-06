package handler

import (
	"net/http"
	"encoding/json"
	"go-web/internal/products/model"
	"github.com/go-chi/chi/v5"
	"go-web/pkg/errors"
	"strconv"
	"os"
)

func (h *ProductHandler) HandlerUpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	token := r.Header.Get("API_TOKEN")
    if token != os.Getenv("API_TOKEN") {
        e := errors.ErrUnauthorized
        e.WriteResponse(w, nil)
        return
    }
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		e := errors.CreateError(http.StatusBadRequest, "impossible conversion of id to int")
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: nil,
			Error: true,
		}
		e.WriteResponse(w, responseBody)
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		e := errors.ErrProductNotFound
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: nil,
			Error: true,
		}
		e.WriteResponse(w, responseBody)
	}
	var requestBody model.RequestBodyProduct
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		e := errors.ErrSaveOrUpdateProduct
	
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: nil,
			Error: true,
		}
		e.WriteResponse(w, responseBody)
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
			e := errors.ErrUpdateProductFailure
			responseBody := model.ResponseBodyProduct{
				Message:	e.Message,
				Product: 	nil,
				Error:	 	true,
			}
			e.WriteResponse(w, responseBody)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
		return
	}
	res, err = h.Service.CreateProduct(product)
	if err != nil {
		e := errors.ErrCreateProductFailure
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product:	nil,
			Error: 		true,
		}
		e.WriteResponse(w, responseBody)
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
	token := r.Header.Get("API_TOKEN")
    if token != os.Getenv("API_TOKEN") {
		e := errors.ErrUnauthorized
        e.WriteResponse(w, nil)
        return
    }
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		e := errors.CreateError(http.StatusBadRequest, "impossible conversion of id to int")
        e.WriteResponse(w, nil)
		return
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		e := errors.ErrProductNotFound
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: nil,
			Error: true,
		}
		e.WriteResponse(w, responseBody)
		return
	}
	if p == nil {
		e := errors.ErrProductNotFound
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: 	nil,
			Error: 		true,
		}
		e.WriteResponse(w, responseBody)
		return
	}
	var requestBody model.RequestUpdateBodyProduct
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		e := errors.ErrUpdateProductFailure
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: 	nil,
			Error: 		true,
		}
		e.WriteResponse(w, responseBody)
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
		e := errors.ErrUpdateProductFailure
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: 	nil,
			Error: 		true,
		}
		e.WriteResponse(w, responseBody)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}