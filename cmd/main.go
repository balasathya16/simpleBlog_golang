package main

import (
	"fmt"
	"log"
	"net/http"

	"simpleBlog/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the Simple Blog")

	router := mux.NewRouter()
	routes.ConfigureRoutes(router)

	// server configurations and routes

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
