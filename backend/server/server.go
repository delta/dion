package server

import (
	"fmt"
	"net/http"
	"time"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/server/routes"
	"delta.nitt.edu/dion/services/logging"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if config.C.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	routes.InitRoutes()
	for group, routeList := range routes.RouteMap {
		rg := router.Group(group)
		rg.Use(routeList.GlobalMiddleware...)
		for _, route := range routeList.Routes {
			var handlerFunc []gin.HandlerFunc
			if route.Middleware == nil {
				handlerFunc = []gin.HandlerFunc{route.HandlerFunc}
			} else {
				handlerFunc = append(route.Middleware, route.HandlerFunc)
			}
			switch route.Method {
			case http.MethodGet:
				rg.GET(route.Pattern, handlerFunc...)
			case http.MethodPost:
				rg.POST(route.Pattern, handlerFunc...)
			case http.MethodPut:
				rg.PUT(route.Pattern, handlerFunc...)
			case http.MethodPatch:
				rg.PATCH(route.Pattern, handlerFunc...)
			case http.MethodDelete:
				rg.DELETE(route.Pattern, handlerFunc...)
			}
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
