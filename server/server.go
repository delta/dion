package server

import (
	"context"

	"delta.nitt.edu/dion/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Params struct {
	fx.In
	Config *config.Config
}

type RouterParams struct {
	Params
	R  *gin.Engine
	Db *gorm.DB
}

func newLogger(p Params) (*zap.Logger, error) {
	if p.Config.IsProd {
		return zap.NewProduction()
	} else {
		return zap.NewDevelopment()
	}
}

func NewServer(lc fx.Lifecycle, g *gin.Engine, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Sugar().Infow("Starting Server")
			return g.Run(":8000")
		},
		OnStop: func(ctx context.Context) error {
			logger.Sugar().Infow("Stopping Server")
			return nil
		},
	})
}

func ginInstance() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	return r
}

var Module = fx.Options(
	fx.Provide(
		ginInstance,
		newLogger,
	),
)
