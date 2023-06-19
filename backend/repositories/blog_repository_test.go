package repositories

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"simpleBlog/backend/models"
)

func TestBlogRepository_Create(t *testing.T) {
	repo := NewBlogRepository()

	// Generate a new ObjectID
	id := primitive.NewObjectID()

	blog := &models.Blog{
		ID:      id,
		Title:   "My First Blog Post",
		Content: "Hello, world!",
		Author:  "John Doe",
	}

	err := repo.Create(blog)
	assert.NoError(t, err)

	// Additional assertions or checks can be performed here

	// Retrieve the blog post by ID
	fmt.Println("Getting blog by ID:", blog.ID.Hex())
	createdBlog, err := repo.GetByID(blog.ID.Hex())
	if err != nil {
		log.Fatal("Error retrieving blog by ID:", err)
	}
	fmt.Println("Retrieved blog by ID:", createdBlog)

	assert.Equal(t, blog.Title, createdBlog.Title)
	assert.Equal(t, blog.Content, createdBlog.Content)
	assert.Equal(t, blog.Author, createdBlog.Author)

	// Update the blog post
	blog.Content = "Updated content"
	err = repo.Update(blog)
	assert.NoError(t, err)

	// Retrieve the blog post again to verify the update
	fmt.Println("Getting updated blog by ID:", blog.ID.Hex())
	updatedBlog, err := repo.GetByID(blog.ID.Hex())
	if err != nil {
		log.Fatal("Error retrieving updated blog by ID:", err)
	}
	fmt.Println("Retrieved updated blog by ID:", updatedBlog)

	assert.Equal(t, blog.Content, updatedBlog.Content)

	// Delete the blog post
	err = repo.Delete(blog.ID.Hex())
	assert.NoError(t, err)

	// Try to retrieve the deleted blog post
	fmt.Println("Getting deleted blog by ID:", blog.ID.Hex())
	deletedBlog, err := repo.GetByID(blog.ID.Hex())
	if err == nil {
		log.Fatal("Expected error retrieving deleted blog by ID, but got no error.")
	}
	fmt.Println("Deleted blog:", deletedBlog)

	assert.Error(t, err)
	assert.Nil(t, deletedBlog)
}

// ...
