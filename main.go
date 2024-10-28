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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server is running")
}
func DatabaseSQL() {
	err := sql.ConnectSql()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
}
func server(config *config.Config) {
	DatabaseSQL()
	http.HandleFunc("/", index)
	handler := routes.Routes()
	err := http.ListenAndServe(":"+config.Port, middlewares.CORS(handler))
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	config := config.NewConfig()
	server(config)
}
