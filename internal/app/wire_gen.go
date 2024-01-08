// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"net/http"
	"simple-to-do/internal/config/client"
	"simple-to-do/internal/repositories"
	"simple-to-do/internal/services"
	"simple-to-do/internal/transport/handler"
	"simple-to-do/internal/transport/routes"
	"simple-to-do/pkg/validator"
)

// Injectors from wire.go:

func InitializeApp() (*http.Server, error) {
	repositoriesRepositories := repositories.InitalizeTodoDatamaster()
	db, err := client.InitializedDatamaster()
	if err != nil {
		return nil, err
	}
	validate := validator.NewValidator()
	service := services.InitalizeTodoService(repositoriesRepositories, db, validate)
	handlerHandler := handler.InitalizedTodoHandler(service)
	server := routes.InitalizeServer(handlerHandler)
	return server, nil
}
