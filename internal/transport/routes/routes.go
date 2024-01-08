package routes

import (
	"net/http"
	"simple-to-do/internal/transport/handler"

	"github.com/julienschmidt/httprouter"
)

func setupHandler(td handler.Handler) *httprouter.Router {
	r := httprouter.New()
	InitalizeTodoRouter(r, td)
	return r
}

func InitalizeServer(td handler.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: setupHandler(td),
	}
}
