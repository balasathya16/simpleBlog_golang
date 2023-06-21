package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"simpleBlog/backend/config"
	"simpleBlog/backend/repositories"
)

func main() {
	cfg := config.LoadConfig()

	// Create a new instance of the BlogRepository
	repo := repositories.NewBlogRepository(cfg.MongoDBURI, cfg.DatabaseName)

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define the routes and their corresponding handlers
	router.HandleFunc("/blogs", GetBlogs(repo)).Methods("GET")
	router.HandleFunc("/blogs/{id}", GetBlogByID(repo)).Methods("GET")
	router.HandleFunc("/blogs", CreateBlog(repo)).Methods("POST")
	router.HandleFunc("/blogs/{id}", UpdateBlog(repo)).Methods("PUT")
	router.HandleFunc("/blogs/{id}", DeleteBlog(repo)).Methods("DELETE")

	// Start the HTTP server
	addr := fmt.Sprintf("%s%s", cfg.ServerAddress, router)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}

func GetBlogs(repo *repositories.BlogRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your logic to get blogs using the repository
	}
}

func GetBlogByID(repo *repositories.BlogRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your logic to get a blog by ID using the repository
	}
}

func CreateBlog(repo *repositories.BlogRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your logic to create a blog using the repository
	}
}

func UpdateBlog(repo *repositories.BlogRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your logic to update a blog using the repository
	}
}

func DeleteBlog(repo *repositories.BlogRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your logic to delete a blog using the repository
	}
}
