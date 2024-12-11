package controller

import (
	"net/http"
	"encoding/json"
	"go-web/internal/products/model"
	"github.com/go-chi/chi/v5"
	pkgErr "go-web/pkg/error"
	"strconv"
)

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// token := r.Header.Get("API_TOKEN")
    // if token != os.Getenv("API_TOKEN") {
    //     e := pkgErr.ErrUnauthorized
    //     WriteErrorResponse(w, e)
    //     return
    // }
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		e := pkgErr.CreateError(http.StatusBadRequest, "impossible conversion of id to int")
		WriteErrorResponse(w, e)
		return
	}
	p, err := c.Service.GetProductById(nId)
	if err != nil {
		e := pkgErr.ErrProductNotFound
		WriteErrorResponse(w, e)
		return
	}
	var requestBody RequestBodyProduct
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		e := pkgErr.ErrSaveOrUpdateProduct
		WriteErrorResponse(w, e)
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
	if p != nil {
		updatedProduct, err := c.Service.UpdateProduct(product, p.Id)
		if err != nil {
			e := pkgErr.ErrUpdateProductFailure
			WriteErrorResponse(w, e)
			return
		}
		res := ResponseBodyProduct{
			Message: "product updated",
			Product: updatedProduct,
			Error:   false,
		}
		WriteResponse(w, res, http.StatusOK)
		return
	}
	createdProduct, err := c.Service.CreateProduct(product)
	if err != nil {
		e := pkgErr.ErrCreateProductFailure
		WriteErrorResponse(w, e)
		return
	}
	res := ResponseBodyProduct{
		Message: "product created",
		Product: createdProduct,
		Error:   false,
	}
	WriteResponse(w, res, http.StatusCreated)
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

func (c *ProductController) PatchProduct(w http.ResponseWriter, r *http.Request) {
	// token := r.Header.Get("API_TOKEN")
    // if token != os.Getenv("API_TOKEN") {
	// 	e := pkgErr.ErrUnauthorized
    //     WriteErrorResponse(w, e)
    //     return
    // }
	id := chi.URLParam(r, "id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		e := pkgErr.CreateError(http.StatusBadRequest, "impossible conversion of id to int")
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
	var requestBody RequestUpdateBodyProduct
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		e := pkgErr.ErrUpdateProductFailure
		WriteErrorResponse(w, e)
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
	updatedProduct, err := c.Service.UpdateProduct(product, p.Id)
	if err != nil {
		e := pkgErr.ErrUpdateProductFailure
		WriteErrorResponse(w, e)
		return
	}
	res := ResponseBodyProduct{
		Message: "product updated",
		Product: updatedProduct,
		Error:   false,
	}
	WriteResponse(w, res, http.StatusOK)
	return
}
