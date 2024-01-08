package routes

import (
	"simple-to-do/internal/transport/handler"

	"github.com/julienschmidt/httprouter"
)

func InitalizeTodoRouter(r *httprouter.Router, h handler.Handler) {
	r.GET("/api/tasks", h.All)
	r.GET("/api/tasks/:id", h.Get)
	r.POST("/api/tasks", h.Post)
	r.PUT("/api/tasks/:id", h.Put)
	r.DELETE("/api/tasks/:id", h.Delete)
	r.PUT("/api/tasks/:id/status", h.SetStatus)
}
