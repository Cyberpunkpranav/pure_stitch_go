package routes

import (
	"net/http"
	"pure_stitch/controllers"
	"pure_stitch/middlewares"
)

func Products() http.Handler {
	// Create a new ServeMux
	route := http.NewServeMux()
	route.HandleFunc("/list", middlewares.Http_method(http.MethodGet, controllers.Get_products))
	route.HandleFunc("/media", middlewares.Http_method(http.MethodGet, controllers.Product_images))
	route.HandleFunc("/types/list", middlewares.Http_method(http.MethodGet, controllers.Product_types))
	route.HandleFunc("/types/images", middlewares.Http_method(http.MethodGet, controllers.Product_types_images))
	return route
}
