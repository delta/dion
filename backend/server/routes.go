package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Annonate so that we can generate open-api specification

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Helloo")
}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},
	{
		"Ping",
		http.MethodGet,
		"/ping",
		Ping,
	},
}
