package router

import (
	"fmt"
	"net/http"

	"delta.nitt.edu/dion/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
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

func SetupRouter(p RouterParams) {
	if p.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}
	fmt.Println(p.Db)
	
	p.R.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
