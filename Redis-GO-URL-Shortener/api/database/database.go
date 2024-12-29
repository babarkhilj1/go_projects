package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// Ctx is a global context used for Redis operations.
// It allows Redis commands to be executed with a consistent context.
var Ctx = context.Background()

// CreateClient creates and returns a Redis client connected to the specified database number.
// The Redis connection details are read from environment variables.
func CreateClient(dbNo int) *redis.Client {
	// Initialize a new Redis client with options.
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"), // Redis server address (host:port), loaded from the environment.
		Password: os.Getenv("DB_PASS"), // Redis server password, loaded from the environment (if required).
		DB:       dbNo,                 // The specific Redis database number to connect to.
	})

	// Return the configured Redis client.
	return rdb
}
