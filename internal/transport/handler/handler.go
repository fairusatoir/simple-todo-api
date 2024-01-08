package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	All(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	Get(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	Post(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	Put(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	SetStatus(w http.ResponseWriter, re *http.Request, p httprouter.Params)
}
