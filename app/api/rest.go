package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Rest interface {
	FindAllItems(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	FindItem(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	CreateItem(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	UpdateItem(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	DeleteItem(w http.ResponseWriter, re *http.Request, p httprouter.Params)
	ComplatedItem(w http.ResponseWriter, re *http.Request, p httprouter.Params)
}
