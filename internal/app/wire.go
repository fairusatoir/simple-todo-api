//go:build wireinject
// +build wireinject

package app

import (
	"net/http"
	"simple-to-do/internal/config/client"
	"simple-to-do/internal/repositories"
	"simple-to-do/internal/services"
	"simple-to-do/internal/transport/handler"
	"simple-to-do/internal/transport/routes"
	"simple-to-do/pkg/validator"

	"github.com/google/wire"
)

func InitializeApp() (*http.Server, error) {
	wire.Build(
		repositories.InitalizeTodoDatamaster,
		validator.NewValidator,
		client.InitializedDatamaster,
		services.InitalizeTodoService,
		handler.InitalizedTodoHandler,
		routes.InitalizeServer,
	)

	return nil, nil
}
