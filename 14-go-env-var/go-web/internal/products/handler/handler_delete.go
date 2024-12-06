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

func (h *ProductHandler) HandlerDeleteProduct(w http.ResponseWriter, r *http.Request) {
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
		e :=  errors.CreateError(http.StatusBadRequest,  "impossible conversion of id to int")
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: 	nil,
			Error: 		true,
		}
		e.WriteResponse(w, responseBody)
		return
	}
	p, err := h.Service.GetProductById(nId)
	if err != nil {
		e := errors.ErrProductNotFound
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: 	nil,
			Error: 		true,
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
	var res *model.ResponseBodyProduct
	res, err = h.Service.DeleteProduct(p.Id)
	if err != nil {
		e := errors.ErrDeleteProductFailure
		responseBody := model.ResponseBodyProduct{
			Message:	e.Message,
			Product: nil,
			Error: true,
		}
		e.WriteResponse(w, responseBody)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}