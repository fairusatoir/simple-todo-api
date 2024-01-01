package main

import (
	"net/http"
	"simple-to-do/app/controllers"
	"simple-to-do/utilities"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: controllers.Handler(),
	}

	err := server.ListenAndServe()
	utilities.PanicOnError(err)
}
