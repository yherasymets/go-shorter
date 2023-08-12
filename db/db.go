package db

import (
	"os"

	"github.com/go-redis/redis"
)

func CreateClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
	})

	return rdb
}
