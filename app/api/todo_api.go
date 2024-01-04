package api

import (
	"net/http"
	"simple-to-do/app/domains"
	"simple-to-do/app/usecases"
	"simple-to-do/utilities"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoApi struct {
	Usecase usecases.Usecase
}

func NewApi(u usecases.Usecase) Rest {
	return &TodoApi{
		Usecase: u,
	}
}

func (a *TodoApi) FindAllItems(w http.ResponseWriter, re *http.Request, _ httprouter.Params) {
	items, err := a.Usecase.GetItems(re.Context())
	if err != nil {
		panic(err)
	}
	utilities.GenerateResponse(w, http.StatusOK, items, nil)
}

func (a *TodoApi) FindItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	utilities.PanicOnError(err)

	item, err := a.Usecase.GetItemById(re.Context(), id)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusOK, item, nil)
}

func (a *TodoApi) CreateItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	reqItem := domains.Task{}
	utilities.ReadRequest(re, &reqItem)

	item, err := a.Usecase.InsertItem(re.Context(), reqItem)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusCreated, item, nil)
}

func (a *TodoApi) UpdateItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	utilities.PanicOnError(err)

	reqItem := domains.Task{}
	utilities.ReadRequest(re, &reqItem)

	reqItem.Id = id

	item, err := a.Usecase.UpdateItem(re.Context(), reqItem)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusAccepted, item, nil)
}

func (a *TodoApi) DeleteItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	utilities.PanicOnError(err)

	err = a.Usecase.DeleteItem(re.Context(), id)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusOK, nil, nil)
}

func (a *TodoApi) ComplatedItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	utilities.PanicOnError(err)

	reqItem := domains.UpdateStatusTask{}
	utilities.ReadRequest(re, &reqItem)

	reqItem.Id = id

	item, err := a.Usecase.UpdateCompletedItem(re.Context(), reqItem)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusOK, item, nil)
}
