package routes

import (
	"net/http"
	"pure_stitch/controllers"
	"pure_stitch/middlewares"
)

func Arrivals() http.Handler {
	// Create a new ServeMux
	route := http.NewServeMux()
	route.HandleFunc("/list", middlewares.Http_method(http.MethodGet, controllers.New_arrivals))
	return route
}
