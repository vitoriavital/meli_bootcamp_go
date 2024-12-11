package controller

import (
	"encoding/json"
	"go-web/internal/products/model"
	pkgErr "go-web/pkg/error"
	"net/http"
)

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// token := r.Header.Get("API_TOKEN")
    // if token != os.Getenv("API_TOKEN") {
	// 	e := pkgErr.ErrUnauthorized
	// 	WriteErrorResponse(w, e)
    //     return
    // }
	var requestBody RequestBodyProduct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		e := pkgErr.ErrCreateProductFailure
		WriteErrorResponse(w, e)
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

	p, err := c.Service.CreateProduct(newProduct)
	if err != nil {
		e := pkgErr.ErrCreateProductFailure
		WriteErrorResponse(w, e)
		return
	}
	res := ResponseBodyProduct{
		Message: "new product created",
		Product: p,
		Error: false,
	}
	WriteResponse(w, res, http.StatusCreated)
}
