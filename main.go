package main

import (
	"net/http"
	"simple-to-do/app/api"
	"simple-to-do/app/repositories"
	"simple-to-do/app/usecases"
	"simple-to-do/config"
	"simple-to-do/utilities"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	d := config.InitMysqlMasterData()
	v := validator.New()

	repo := repositories.NewDatamaster()
	usecase := usecases.NewTodoUsecase(repo, d, v)
	router := api.NewApi(usecase)
	handler := config.Router(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	utilities.PanicOnError(err)
}
