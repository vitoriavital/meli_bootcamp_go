package handler

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-web/pkg/errors"
	"strconv"
)

func (h *ProductHandler) HandlerProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	products, err := h.Service.GetAllProducts()
	if err != nil {
		e := errors.ErrUnauthorized
		e.WriteResponse(w, nil)
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
		e := errors.ErrProductNotFound
		e.WriteResponse(w, nil)
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
		e := errors.CreateError(http.StatusBadRequest, "missing priceGt")
        e.WriteResponse(w, nil)
        return
	}
	priceGt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		e := errors.CreateError(http.StatusBadRequest, "not a valid value for priceGt")
        e.WriteResponse(w, nil)
        return
	}
	products,err := h.Service.GetProductsByPrice(priceGt)
	if err != nil {
		e := errors.CreateError(http.StatusInternalServerError, "no products found for this priceGt")
		e.WriteResponse(w, nil)
        return
	}
	w.WriteHeader(http.StatusOK)
	var msg string
	for _, p := range products {
		msg += fmt.Sprint("id: ", p.Id , " - ", p.Name, " - $", p.Price, " - ", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}