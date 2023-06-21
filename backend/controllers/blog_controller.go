package controllers

import (
	"encoding/json"
	"net/http"

	"simpleBlog/backend/repositories"

	"github.com/gorilla/mux"
)

type BlogController struct {
	repo *repositories.BlogRepository
}

func NewBlogController(repo *repositories.BlogRepository) *BlogController {
	return &BlogController{
		repo: repo,
	}
}

func (c *BlogController) GetBlogs(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all blogs
	blogs, err := c.repo.GetAll()
	if err != nil {
		// Handle the error, e.g., return an appropriate HTTP response
	}

	// Send the blogs as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func (c *BlogController) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	// Extract the blog ID from the request parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Call the repository method to get the blog by ID
	blog, err := c.repo.GetByID(id)
	if err != nil {
		// Handle the error (e.g., return an error response)
		http.Error(w, "Failed to get blog by ID", http.StatusInternalServerError)
		return
	}

	// Convert the blog to JSON
	blogJSON, err := json.Marshal(blog)
	if err != nil {
		// Handle the error (e.g., return an error response)
		http.Error(w, "Failed to marshal blog to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type and write the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.Write(blogJSON)
}

func (c *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a new blog
}

func (c *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an existing blog
}

func (c *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a blog
}
