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

//	func index(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("server is running")
//	}
func DatabaseSQL() {
	err := sql.ConnectSql()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
}
func file_handler() http.Handler {
	fileServe := http.FileServer(http.Dir("./assets/arrivals/general"))
	return http.StripPrefix("/arrivals/", fileServe)
}
func server(config *config.Config) {
	DatabaseSQL()
	handler := routes.Routes()
	http.Handle("/arrivals/", file_handler())
	finalHandler := middlewares.CORS(handler)

	err := http.ListenAndServe(":"+config.Port, finalHandler)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	config := config.NewConfig()
	server(config)
	file_handler()
}
