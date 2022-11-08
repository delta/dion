package router

import (
	"net/http"

	"delta.nitt.edu/dion/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	
	if config.C.Environment == "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
