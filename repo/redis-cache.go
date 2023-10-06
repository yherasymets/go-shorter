package repo

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       0,
	})
}
