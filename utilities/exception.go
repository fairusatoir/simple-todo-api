package utilities

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func PanicOnError(err error) {
	if err != nil {
		fmt.Printf("[%s][%s]\n", "ERROR", err)
		panic(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	fmt.Printf("[%s][%s]\n", "ERROR", err)

	if ex, ok := err.(validator.ValidationErrors); ok {
		badRequestError(w, r, ex.Error())
		return
	}

	if ex, ok := err.(NotFoundError); ok {
		notFoundError(w, r, ex.Error)
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	GenerateResponse(w, status, nil, err)
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) {
	status := http.StatusBadRequest
	w.WriteHeader(status)
	GenerateResponse(w, status, nil, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	GenerateResponse(w, status, nil, err)
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(data interface{}) NotFoundError {
	fmt.Printf("[%s][%v]\n", "NOT FOUND", data)
	return NotFoundError{Error: http.StatusText(http.StatusNotFound)}
}
