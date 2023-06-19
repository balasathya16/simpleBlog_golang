package main

import (
	"fmt"
	"log"
	"net/http"

	"simpleBlog/backend/config"
	//"simpleBlog/backend/controllers"
	"simpleBlog/backend/repositories"

	"github.com/gorilla/mux"
	//"simpleBlog/backend/routes"
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
	addr := fmt.Sprintf("%s%s", cfg.ServerAddress, router)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
