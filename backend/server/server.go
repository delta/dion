package server

import (
	"fmt"
	"net/http"
	"time"

	"delta.nitt.edu/dion/utils"

	"github.com/gin-contrib/cors"

	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

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
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	expiry := utils.GetInt("JWT_EXPIRY", 24)
	var secure bool
	env := os.Getenv("ENVIRONMENT")
	if env == "prod" {
		secure = true
	} else {
		secure = false
	}
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   int(time.Duration(expiry) * time.Hour),
		Secure:   secure,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("dashboardAuth", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	routes.InitRoutes()
	for group, routeList := range routes.RouteMap {
		rg := router.Group(group)
		rg.Use(routeList.GlobalMiddleware...)
		for _, route := range routeList.Routes {
			var handlerFunc gin.HandlersChain
			if route.Middleware == nil {
				handlerFunc = gin.HandlersChain{route.HandlerFunc}
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
