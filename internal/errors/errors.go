package errors

import "net/http"

type RestAPIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestAPIError {
	return &RestAPIError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

func NewNotFoundError(message string) *RestAPIError {
	return &RestAPIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestAPIError {
	return &RestAPIError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}
