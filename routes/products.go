package routes

import (
	"net/http"
	"pure_stitch/controllers"
	"pure_stitch/middlewares"
)

func Routes() http.Handler {
	// Create a new ServeMux
	route := http.NewServeMux()
	route.HandleFunc("/products/list", middlewares.Http_method(http.MethodGet, controllers.Get_products))
	route.HandleFunc("/product/types/list", middlewares.Http_method(http.MethodGet, controllers.Product_types))

	return route
}
