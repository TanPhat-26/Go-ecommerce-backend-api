package cache

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(config RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
}
