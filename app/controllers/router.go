package controllers

import (
	"net/http"
	"simple-to-do/app/domains"
	"simple-to-do/app/repositories"
	"simple-to-do/app/usecases"
	"simple-to-do/config"
	"simple-to-do/utilities"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type router struct {
	Usecase usecases.Usecase
}

func NewRouter(u usecases.Usecase) *router {
	return &router{
		Usecase: u,
	}
}

func Handler() *httprouter.Router {
	d := config.InitMysqlMasterData()
	v := validator.New()

	repo := repositories.NewRepositories()
	usecase := usecases.NewUsecase(repo, d, v)
	handler := NewRouter(usecase)

	r := httprouter.New()

	r.GET("/api/tasks", handler.FindAllItems)
	r.GET("/api/tasks/:id", handler.FindItem)
	r.POST("/api/tasks", handler.CreateItem)
	r.PUT("/api/tasks/:id", handler.UpdateItem)

	r.PanicHandler = utilities.ErrorHandler
	return r
}

func (r *router) FindAllItems(w http.ResponseWriter, re *http.Request, _ httprouter.Params) {
	items, err := r.Usecase.GetItems(re.Context())
	if err != nil {
		panic(err)
	}
	utilities.GenerateResponse(w, http.StatusOK, items, nil)
}

func (r *router) FindItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	utilities.PanicOnError(err)

	item, err := r.Usecase.GetItemById(re.Context(), id)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusOK, item, nil)
}

func (r *router) CreateItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	reqItem := domains.Task{}
	utilities.ReadRequest(re, &reqItem)

	item, err := r.Usecase.InsertItem(re.Context(), reqItem)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusCreated, item, nil)
}

func (r *router) UpdateItem(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	utilities.PanicOnError(err)

	reqItem := domains.Task{}
	utilities.ReadRequest(re, &reqItem)

	reqItem.Id = id

	item, err := r.Usecase.UpdateItem(re.Context(), reqItem)
	utilities.PanicOnError(err)

	utilities.GenerateResponse(w, http.StatusAccepted, item, nil)
}
