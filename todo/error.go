package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

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
	GenerateResponse(w, httpRes(status, nil, err))
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) {
	status := http.StatusBadRequest
	w.WriteHeader(status)
	GenerateResponse(w, httpRes(status, nil, err))
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	GenerateResponse(w, httpRes(status, nil, err))
}

func PanicIfError(err error) {
	if err != nil {
		fmt.Printf("[ERROR][%s]\n", err)
		panic(err)
	}
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(id int) NotFoundError {
	fmt.Printf("[NOT FOUND][%s]\n", strconv.Itoa(id))
	return NotFoundError{Error: http.StatusText(http.StatusNotFound)}
}
