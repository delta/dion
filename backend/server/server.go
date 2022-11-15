package server

import (
	"fmt"
	"net/http"
	"time"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/services/logging"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if config.C.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

func StartServer() {
	router := InitRouter()
	maxHeaderBytes := 1 << 20

	// errorLoggerZap := logging.Sugared().Named("gin error:")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.C.Server.Port),
		Handler:        router,
		ReadTimeout:    (time.Duration(config.C.Server.ReadTimeout)) * (time.Second),
		WriteTimeout:   (time.Duration(config.C.Server.WriteTimeout)) * (time.Second),
		MaxHeaderBytes: maxHeaderBytes,
		ErrorLog:       nil,
	}
	logging.Sugared().Infof("Starting the server and listening on port %d", config.C.Server.Port)
	s.ListenAndServe()
}
