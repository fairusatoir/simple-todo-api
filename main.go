package main

import (
	"fairusatoir/simple-to-do/todo"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: todo.SetRouter(),
	}

	err := server.ListenAndServe()
	todo.PanicIfError(err)
}
