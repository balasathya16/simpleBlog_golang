package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting the Simple Blog")

	// server configurations and routes

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
