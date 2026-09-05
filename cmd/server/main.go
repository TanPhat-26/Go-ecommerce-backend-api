package main

import (
	"context"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/initialize"
	"go.uber.org/zap"
)

func main() {
	initialize.LoadConfig()
	initialize.InitLogger()
	defer global.Logger.Sync()

	if err := initialize.InitDatabase(); err != nil {
		global.Logger.Fatal("failed to initialize database", zap.Error(err))
	}
	global.Logger.Info("database connected successfully")

	if err := initialize.SeedRoles(context.Background(), global.DB); err != nil {
		global.Logger.Fatal("failed to seed roles", zap.Error(err))
	}

	if err := initialize.InitRedis(); err != nil {
		global.Logger.Fatal("failed to initialize redis", zap.Error(err))
	}
	global.Logger.Info("redis connected successfully")

	r := initialize.InitRouter()
	addr := ":" + global.Config.App.Port

	global.Logger.Info(
		"starting http server",
		zap.String("addr", addr),
		zap.String("app", global.Config.App.Name),
		zap.String("env", global.Config.App.Env),
	)

	if err := r.Run(addr); err != nil {
		global.Logger.Fatal("failed to start http server", zap.Error(err))
	}
}
