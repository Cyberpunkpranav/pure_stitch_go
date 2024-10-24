package main

import (
	"fmt"
	"log"
	"net/http"
	"pure_stitch/config"
	sql "pure_stitch/sql"
	_ "github.com/go-sql-driver/mysql"

)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server is running")
}
func DatabaseSQL(){
	err:=sql.ConnectSql()
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

}
func server(config *config.Config) {
	DatabaseSQL()
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	config := config.NewConfig()
	server(config)
}
