package main

import (
	"fmt"
	"log"
	"net/http"
	"pure_stitch/config"
	"pure_stitch/middlewares"
	"pure_stitch/routes"
	sql "pure_stitch/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseSQL() {
	err := sql.ConnectSql()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
}

func server(config *config.Config) {
	//Database connection

	DatabaseSQL()

	// APi's routing

	// products
	products := routes.Products()

	productsHandler := http.StripPrefix("/api/products", products)

	http.Handle("/api/products/", middlewares.CORS(productsHandler))

	// new-arrivals images

	// images := routes.Images()

	// imagesHandler := http.StripPrefix("/assets/arrivals-images", images)

	// http.Handle("/assets/arrivals-images/", middlewares.CORS(imagesHandler))

	fmt.Println("Server running on http://localhost:" + config.Port)

	err := http.ListenAndServe(":"+config.Port, nil)

	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	config := config.NewConfig()
	server(config)
}
