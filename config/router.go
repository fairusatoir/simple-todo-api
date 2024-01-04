package config

import (
	"simple-to-do/app/api"
	"simple-to-do/utilities"

	"github.com/julienschmidt/httprouter"
)

func Router(a api.Rest) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tasks", a.FindAllItems)
	router.GET("/api/tasks/:id", a.FindItem)
	router.POST("/api/tasks", a.CreateItem)
	router.PUT("/api/tasks/:id", a.UpdateItem)
	router.DELETE("/api/tasks/:id", a.DeleteItem)
	router.PUT("/api/tasks/:id/status", a.ComplatedItem)

	router.PanicHandler = utilities.ErrorHandler
	return router
}
