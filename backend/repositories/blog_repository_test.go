package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"simpleBlog/backend/models"
)

func TestBlogRepository_Create(t *testing.T) {
	repo := NewBlogRepository()

	blog := &models.Blog{
		Title:   "My First Blog Post",
		Content: "Hello, world!",
		Author:  "John Doe",
	}

	err := repo.Create(blog)
	assert.NoError(t, err)

	// Additional assertions or checks can be performed here

	// Retrieve the blog post by ID
	createdBlog, err := repo.GetByID(blog.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, blog.Title, createdBlog.Title)
	assert.Equal(t, blog.Content, createdBlog.Content)
	assert.Equal(t, blog.Author, createdBlog.Author)

	// Update the blog post
	blog.Content = "Updated content"
	err = repo.Update(blog)
	assert.NoError(t, err)

	// Retrieve the blog post again to verify the update
	updatedBlog, err := repo.GetByID(blog.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, blog.Content, updatedBlog.Content)

	// Delete the blog post
	err = repo.Delete(blog.ID.Hex())
	assert.NoError(t, err)

	// Try to retrieve the deleted blog post
	deletedBlog, err := repo.GetByID(blog.ID.Hex())
	assert.Error(t, err)
	assert.Nil(t, deletedBlog)
}

// ...
