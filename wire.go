//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"simple-to-do/app/api"
	"simple-to-do/app/repositories"
	"simple-to-do/app/usecases"
	"simple-to-do/config"

	"github.com/google/wire"
)

func InitializedApp() *http.Server {
	wire.Build(
		config.InitMysqlMasterData,
		config.NewValidator,
		repositories.NewDatamaster,
		usecases.NewTodoUsecase,
		api.NewApi,
		config.InitHandler,
	)

	return nil
}
