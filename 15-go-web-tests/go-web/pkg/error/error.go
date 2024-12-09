package error

import (
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
	ErrCreateProductFailure  = CreateError(http.StatusBadRequest, "couldn't create product")
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
