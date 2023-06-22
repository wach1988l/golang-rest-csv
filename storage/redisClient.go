package storage

import (
	"github.com/go-redis/redis/v8"
)

const fileVersion = "file_version"

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.Client.Context(), key).Result()
}
