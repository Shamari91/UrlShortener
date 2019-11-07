package infrastructure

import "github.com/go-redis/redis"

type RedisHandler struct {
	Client *redis.Client
}

func (redisHandler *RedisHandler) Set(key, value string) error {
	return redisHandler.Client.Set(key, value, 0).Err()
}

func (redisHandler *RedisHandler) Get(key string) (string, error) {
	return redisHandler.Client.Get(key).Result()
}

func NewRedisClient() *RedisHandler {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redisHandler := new(RedisHandler)
	redisHandler.Client = client
	return redisHandler
}
