package router

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"delta.nitt.edu/dion/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Params struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
}

type RouterParams struct {
	Params
	R  *gin.Engine
	Db *gorm.DB
}

func SetupRouter(p RouterParams) {
	if p.Config.Environment == "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	fmt.Println(p.Db)
	
	p.R.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	logger := p.Logger.Sugar()
	configurationsGroup := p.R.Group("/configurations")
	addConfigurationRoutes(configurationsGroup, p.Db, logger)

	projectsGroup := p.R.Group("/projects")
	addProjectRoutes(projectsGroup, p.Db, logger)
}
