package main

import (
	"delta.nitt.edu/dion/server"
	"delta.nitt.edu/dion/services/logging"
	"go.uber.org/zap"
)

func main() {
	logging.Setup(
		logging.WrapOptions(
			zap.AddCaller(),
		))
	defer logging.Flush()
	server.StartServer()
}
