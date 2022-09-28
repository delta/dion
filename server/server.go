package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func newLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

var Module = fx.Provide(
	gin.Default,
	newLogger,
)
