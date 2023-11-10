package repo

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	addr = os.Getenv("DB_ADDR")
	pass = os.Getenv("DB_PASS")
)

func RedisCache() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}
