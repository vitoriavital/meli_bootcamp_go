package handler

import (
	"encoding/json"
	"go-web/internal/products/model"
	"go-web/pkg/errors"
	"net/http"
	"os"
)

func (h *ProductHandler) HandlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	token := r.Header.Get("API_TOKEN")
    if token != os.Getenv("API_TOKEN") {
		e := errors.ErrUnauthorized
		e.WriteResponse(w)
        return
    }
	var requestBody model.RequestBodyProduct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		e := errors.ErrCreateProductFailure
		e.WriteResponse(w)
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
		e := errors.ErrCreateProductFailure
		e.WriteResponse(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}