package errors

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	ErrProductNotFound       = CreateError(http.StatusNotFound, "product not found")
	ErrBadRequest            = CreateError(http.StatusBadRequest, "bad request")
	ErrInternal              = CreateError(http.StatusInternalServerError, "internal server error")
	ErrUnauthorized          = CreateError(http.StatusUnauthorized, "unauthorized")
	ErrCreateProductFailure  = CreateError(http.StatusInternalServerError, "couldn't create product")
	ErrUpdateProductFailure  = CreateError(http.StatusInternalServerError, "couldn't update product")
	ErrDeleteProductFailure  = CreateError(http.StatusInternalServerError, "couldn't delete product")
	ErrSaveOrUpdateProduct   = CreateError(http.StatusInternalServerError, "couldn't update or create product")
)

func CreateError(status int, msg string) CustomError {
	return CustomError{
		Status:  status,
		Message: msg,
	}
}

func (e *CustomError) WriteResponse(w http.ResponseWriter, resBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	err := json.NewEncoder(w).Encode(e)
	if resBody != nil {
		err = json.NewEncoder(w).Encode(resBody)
	}
	if err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}