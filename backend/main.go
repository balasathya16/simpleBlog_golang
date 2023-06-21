package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"simpleBlog/backend/config"
	"simpleBlog/backend/controllers"
	"simpleBlog/backend/repositories"
	"simpleBlog/backend/routes"
)

func main() {
	cfg := config.LoadConfig()

	// Create a new instance of the BlogRepository
	repo := repositories.NewBlogRepository()

	// Create a new instance of the BlogController and pass the repository
	controller := controllers.NewBlogController(repo)

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Register the blog routes
	routes.RegisterBlogRoutes(router, controller)

	// Start the HTTP server
	addr := cfg.ServerAddress
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
