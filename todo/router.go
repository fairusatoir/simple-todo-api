package todo

import (
	"fairusatoir/simple-to-do/todo/domain"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func FindAllItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items := ListAll(r.Context())
	GenerateResponse(w, httpRes(http.StatusOK, items, nil))
}

func FindItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	PanicIfError(err)

	items := GetItemById(r.Context(), id)
	GenerateResponse(w, httpRes(http.StatusOK, items, nil))
}

func CreateItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	item := domain.Task{}
	ReadRequest(r, &item)

	itemResponse := SaveItem(r.Context(), item)
	GenerateResponse(w, httpRes(http.StatusOK, itemResponse, nil))
}

func SetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/tasks", FindAllItem)
	router.GET("/api/tasks/:id", FindItem)
	router.POST("/api/tasks", CreateItem)

	router.PanicHandler = ErrorHandler

	return router
}
