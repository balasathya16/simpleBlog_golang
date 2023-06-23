package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"simpleBlog/backend/config"
	"simpleBlog/backend/controllers"
	"simpleBlog/backend/repositories"
	"simpleBlog/backend/routes"

	"github.com/gorilla/handlers"
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

	// Apply CORS middleware
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Add the origin of your frontend application
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	c := cors(router)

	// Start the HTTP server
	addr := cfg.ServerAddress
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, c))
}
