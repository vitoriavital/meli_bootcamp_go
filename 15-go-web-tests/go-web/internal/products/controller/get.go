package controller

import (
	"go-web/internal/products/model"
	pkgErr "go-web/pkg/error"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	products, err := c.Service.GetAllProducts()
	if err != nil {
		e := pkgErr.CreateError(http.StatusBadRequest, "failed to get all products")
		WriteErrorResponse(w, e)
        return
	}
	var list []model.Product
	for _, p := range products {
		list = append(list, p)
	}
	res := ResponseGetBodyProducts{
		Message: "success to get all products",
		Products: &list,
		Error: false,
	}
	WriteResponse(w, res, http.StatusOK)
}

func (c *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		e := pkgErr.CreateError(http.StatusBadRequest, "invalid value for id")
		WriteErrorResponse(w, e)
		return
	}
	p, err := c.Service.GetProductById(nId)
	if err != nil || p == nil {
		e := pkgErr.ErrProductNotFound
		WriteErrorResponse(w, e)
        return
	}
	res := ResponseBodyProduct{
		Message: "product found",
		Product: p,
		Error: false,
	}

	WriteResponse(w, res, http.StatusOK)
}

func (c *ProductController) SearchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	price := r.URL.Query().Get("priceGt")
	if price == "" {
		e := pkgErr.CreateError(http.StatusBadRequest, "missing priceGt")
        WriteErrorResponse(w, e)
        return
	}
	priceGt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		e := pkgErr.CreateError(http.StatusBadRequest, "not a valid value for priceGt")
        WriteErrorResponse(w, e)
        return
	}
	products,err := c.Service.GetProductsByPrice(priceGt)
	if err != nil || len(products) < 1 {
		e := pkgErr.CreateError(http.StatusInternalServerError, "no products found for this priceGt")
		WriteErrorResponse(w, e)
        return
	}
	var list []model.Product
	for _, p := range products {
		list = append(list, p)
	}
	res := ResponseGetBodyProducts{
		Message: "products found",
		Products: &list,
		Error: false,
	}
	WriteResponse(w, res, http.StatusOK)
}