package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to handle the home page request
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Welcome to the home page!")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to handle the request for individual blog posts
}

func ConfigureRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/posts/{id}", PostHandler).Methods("GET")
	// Add more routes as needed
}
