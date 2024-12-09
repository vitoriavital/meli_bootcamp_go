package controller

import (
	"encoding/json"
	"net/http"
	pkgErr "go-web/pkg/error"
	"go-web/internal/products/model"
)


type RequestBodyProduct struct {
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code_value"`
	IsPublished	bool	`json:"is_published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

type ResponseBodyProduct struct {
	Message		string			`json:"message"`
	Product		*model.Product 	`json:"product,omitempty"`
	Error		bool			`json:"error"`
}

type ResponseGetBodyProducts struct {
	Message		string			 		`json:"message"`
	Products	*[]model.Product 		`json:"products,omitempty"`
	Error		bool					`json:"error"`
}

type RequestUpdateBodyProduct struct {
	Name		*string		`json:"name,omitempty"`
	Quantity	*int		`json:"quantity,omitempty"`
	CodeValue	*string		`json:"code_value,omitempty"`
	IsPublished	*bool		`json:"is_published,omitempty"`
	Expiration	*string		`json:"expiration,omitempty"`
	Price		*float64	`json:"price,omitempty"`
}

func WriteErrorResponse(w http.ResponseWriter, e pkgErr.CustomError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	resErr := ResponseBodyProduct{
		Message: e.Message,
		Product: nil,
		Error: true,
	}
	err := json.NewEncoder(w).Encode(resErr)
	if err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}

func WriteResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
