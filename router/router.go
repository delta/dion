package router

import (
	"net/http"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/project"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
	if p.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	p.R.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	projectGroup := p.R.Group("/project")
	projectGroup.POST("new", func(c *gin.Context) {
		project.AddProject(p.Db, c, p.Params.Logger)
	})
	projectGroup.GET("all", func(c *gin.Context) {
		project.GetAllProjects(p.Db, c, p.Params.Logger)
	})
	projectGroup.DELETE("delete", func(c *gin.Context) {
		project.DeleteProject(p.Db, c, p.Params.Logger)
	})
	projectGroup.PUT("all", func(c *gin.Context) {
		project.ChangeProject(p.Db, c, p.Params.Logger)
	})
}
