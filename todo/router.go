package todo

import (
	"encoding/json"
	"fairusatoir/simple-to-do/todo/domain"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ReadReqBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func FindAllItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items := ListAll(r.Context())
	WriteResBody(w, items)
}

func CreateItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	item := domain.Task{}
	ReadReqBody(r, &item)

	itemResponse := SaveItem(r.Context(), item)
	WriteResBody(w, itemResponse)
}

func SetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/tasks", FindAllItem)
	router.POST("/api/tasks", CreateItem)

	// router.PanicHandler = exce
	return router
}
