package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"simpleBlog/backend/config"
	"simpleBlog/backend/controllers"
	"simpleBlog/backend/repositories"
)

func main() {
	cfg := config.LoadConfig()

	// Create a new instance of the BlogRepository
	repo := repositories.NewBlogRepository()

	// Create a new instance of the BlogController and pass the repository
	controller := controllers.NewBlogController(repo)

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define the routes and their corresponding handlers
	router.HandleFunc("/blogs", controller.GetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id}", controller.GetBlogByID).Methods("GET")
	router.HandleFunc("/blogs", controller.CreateBlog).Methods("POST")
	router.HandleFunc("/blogs/{id}", controller.UpdateBlog).Methods("PUT")
	router.HandleFunc("/blogs/{id}", controller.DeleteBlog).Methods("DELETE")

	// Start the HTTP server
	addr := cfg.ServerAddress
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
