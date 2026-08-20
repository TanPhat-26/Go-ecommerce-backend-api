package main

import (
	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/initialize"
	"go.uber.org/zap"
)

func main() {
	initialize.LoadConfig()
	initialize.InitLogger()

	defer global.Logger.Sync()

	global.Logger.Info(
		"starting e-commerce backend api",
		zap.String("app", global.Config.App.Name),
		zap.String("env", global.Config.App.Env),
		zap.String("port", global.Config.App.Port),
	)
}
