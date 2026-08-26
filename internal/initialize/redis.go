package initialize

import (
	"context"
	"fmt"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/cache"
)

func InitRedis() error {
	client := cache.NewRedisClient(cache.RedisConfig{
		Addr:     global.Config.Redis.Addr(),
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("ping redis: %w", err)
	}

	global.Redis = client
	return nil
}
