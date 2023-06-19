package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config represents the application configuration settings.
type Config struct {
	MongoDBURI    string
	DatabaseName  string
	ServerAddress string
}

// LoadConfig loads the configuration settings from environment variables or other sources.
func LoadConfig() *Config {
	return &Config{
		MongoDBURI:    getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		DatabaseName:  getEnv("DATABASE_NAME", "blogdb"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
	}
}

// getEnv retrieves the value of an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// GetMongoDBURI returns the MongoDB URI from the application configuration.
func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current working directory:", dir)

	err = godotenv.Load(filepath.Join(dir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func GetMongoDBURI() string {
	return os.Getenv("MONGODB_URI")
}

// GetDatabaseName returns the database name from the application configuration.
func GetDatabaseName() string {
	return os.Getenv("DATABASE_NAME")
}

// GetServerAddress returns the server address from the application configuration.
func GetServerAddress() string {
	return appConfig.ServerAddress
}

// appConfig holds the instance of the application configuration.
var appConfig *Config

func init() {
	appConfig = LoadConfig()
}
