package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"simpleBlog/backend/models"
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
	// Parse the request body to extract the blog data
	var blog models.Blog
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the blog data
	if err := validateBlog(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the repository method to create the blog
	err = c.repo.Create(&blog)
	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}

	// Set the response status and content type
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// Send the created blog as the response body
	json.NewEncoder(w).Encode(blog)
}

func validateBlog(blog *models.Blog) error {
	// Implement your validation logic here
	// For example, you can check if the required fields are present
	if blog.Title == "" || blog.Content == "" || blog.Author == "" {
		return errors.New("required fields are missing")
	}

	return nil
}

func (c *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	// Parse the blog ID from the request URL
	vars := mux.Vars(r)
	blogID := vars["id"]

	// Parse the request body to extract the updated blog data
	var updatedBlog models.Blog
	err := json.NewDecoder(r.Body).Decode(&updatedBlog)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the updated blog data
	if err := validateBlog(&updatedBlog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the existing blog from the repository
	existingBlog, err := c.repo.GetByID(blogID)
	if err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	// Update the fields of the existing blog with the updated data
	existingBlog.Title = updatedBlog.Title
	existingBlog.Content = updatedBlog.Content
	existingBlog.Author = updatedBlog.Author

	// Call the repository method to update the blog
	err = c.repo.Update(existingBlog)
	if err != nil {
		http.Error(w, "Failed to update blog", http.StatusInternalServerError)
		return
	}

	// Set the response status and content type
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Send the updated blog as the response body
	json.NewEncoder(w).Encode(existingBlog)
}

func (c *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	// Parse the blog ID from the request URL
	vars := mux.Vars(r)
	blogID := vars["id"]

	// Call the repository method to delete the blog
	err := c.repo.Delete(blogID)
	if err != nil {
		http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
		return
	}

	// Set the response status
	w.WriteHeader(http.StatusOK)
}
