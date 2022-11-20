package main

import (
	"delta.nitt.edu/dion/repository"
	"delta.nitt.edu/dion/server"
	"delta.nitt.edu/dion/services/logging"
	"go.uber.org/zap"
)

func main() {
	repository.Init()
	logging.Setup(
		logging.WrapOptions(
			zap.AddCaller(),
		))
	defer logging.Flush()
	// repository.InsertUser("a", "a@a.com")
	server.StartServer()
}
