package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Helloo")
}

func Pass(ctx *gin.Context) {
	fmt.Println("Hello")
	ctx.Next()
}

func initDefault() {
	group := "default"
	RouteMap[group] = RouteGroup{
		Routes: Routes{
			{
				"Index",
				http.MethodGet,
				"/",
				Index,
				nil,
			},
			{
				"Ping",
				http.MethodGet,
				"/ping",
				Ping,
				nil,
			},
		},
	}
}
