package controllers

import (
	"net/http"

	"simpleBlog/backend/repositories"
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
}

func (c *BlogController) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a blog by ID
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
