package routes

import (
	"net/http"

	"github.com/vallard/trudiedo-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"TodoIndex",
		"GET",
		"/v1/todos",
		handlers.TodoIndex,
	},
}
