package main

import (
	"context"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/controllers"
	"delta.nitt.edu/dion/middlewares"
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/repository"
	"delta.nitt.edu/dion/server"
	"go.uber.org/fx"
)

var App = fx.Options(
	server.Module,
	repository.Module,
	config.Module,
	controllers.Module,
	middlewares.Module,
	models.Module,
	fx.Invoke(
		repository.Migrate,
		server.SetupRouter,
		server.NewServer,
	),
)

func main() {
	fx.New(App).Start(context.Background())
}
