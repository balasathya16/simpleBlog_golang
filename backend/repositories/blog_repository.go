package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"simpleBlog/backend/config"
	"simpleBlog/backend/models"
)

type BlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository() *BlogRepository {
	uri := config.GetMongoDBURI()
	log.Println("MongoDB URI:", uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	database := client.Database(config.GetDatabaseName())
	collection := database.Collection("blogs")

	return &BlogRepository{
		collection: collection,
	}
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(context.TODO(), blog)
	return err
}

// get all blogs

func (r *BlogRepository) GetAll() ([]*models.Blog, error) {
	filter := bson.M{} // Empty filter to retrieve all blogs

	// Define options for sorting or pagination if needed
	options := options.Find()

	// Execute the query and retrieve the blogs
	cur, err := r.collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var blogs []*models.Blog

	// Iterate over the query results and decode each blog
	for cur.Next(context.TODO()) {
		var blog models.Blog
		err := cur.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *BlogRepository) GetByID(id string) (*models.Blog, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var blog models.Blog
	err = r.collection.FindOne(context.TODO(), filter).Decode(&blog)
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (r *BlogRepository) Update(blog *models.Blog) error {
	blog.UpdatedAt = time.Now()

	filter := bson.M{"_id": blog.ID}
	update := bson.M{"$set": blog}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *BlogRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	_, err = r.collection.DeleteOne(context.TODO(), filter)
	return err
}
