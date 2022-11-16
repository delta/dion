package routes

import (
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
	// Middleware function to run
	Middleware []gin.HandlerFunc
}

// Routes is a list of Routes.
type Routes []Route

type RouteGroup struct {
	// List of all routes belonging to the group
	Routes Routes
	// List of middleware to be added to the group
	GlobalMiddleware []gin.HandlerFunc
}

var RouteMap map[string]RouteGroup

func InitRoutes() {
	RouteMap = make(map[string]RouteGroup)
	initDefault()
}
