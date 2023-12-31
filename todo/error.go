package todo

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err string) string {
	fmt.Printf("[ERROR][%s]", err)
	return err
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if ex, ok := err.(validator.ValidationErrors); ok {
		badRequestError(w, r, ex.Error())
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

func PanicIfError(err error) {
	if err != nil {
		panic(ErrorResponse(err.Error()))
	}
}
