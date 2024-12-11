package controller

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	pkgErr "go-web/pkg/error"
	"strconv"
)

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// token := r.Header.Get("API_TOKEN")
    // if token != os.Getenv("API_TOKEN") {
	// 	e := pkgErr.ErrUnauthorized
    //     WriteErrorResponse(w, e)
    //     return
    // }
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		e :=  pkgErr.CreateError(http.StatusBadRequest,  "impossible conversion of id to int")
		WriteErrorResponse(w, e)
		return
	}
	p, err := c.Service.GetProductById(nId)
	if err != nil {
		e := pkgErr.ErrProductNotFound
		WriteErrorResponse(w, e)
		return
	}
	if p == nil {
		e := pkgErr.ErrProductNotFound
		WriteErrorResponse(w, e)
		return
	}
	err = c.Service.DeleteProduct(p.Id)
	if err != nil {
		e := pkgErr.ErrDeleteProductFailure
		WriteErrorResponse(w, e)
		return
	}
	res := ResponseBodyProduct{
		Message: "product deleted",
		Product: nil,
		Error: false,
	}
	WriteResponse(w, res, http.StatusOK)
	return
}
