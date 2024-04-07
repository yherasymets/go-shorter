package repo

import (
	"os"

	"github.com/go-redis/redis/v8"
)

// Retrieve Redis connection parameters from environment variables.
var (
	addr = os.Getenv("DB_ADDR") // Redis server address.
	pass = os.Getenv("DB_PASS") // Redis server password (if any).
)

// NewCache creates a new Redis client instance.
func NewCache() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}
