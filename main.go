package main

import (
	"fmt"
	"log"
	"net/http"
	"pure_stitch/config"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server is running")
}
func server(config *config.Config) {
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
