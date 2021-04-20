package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	restError := RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
	return &restError
}

func NewNotFoundError(message string) *RestErr {
	restError := RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Bad Request",
	}
	return &restError
}