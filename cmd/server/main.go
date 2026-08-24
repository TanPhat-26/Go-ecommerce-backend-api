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

	r := initialize.InitRouter()
	addr := ":" + global.Config.App.Port

	global.Logger.Info(
		"starting http server",
		zap.String("addr", addr),
		zap.String("app", global.Config.App.Name),
		zap.String("env", global.Config.App.Env),
	)

	if err := r.Run(addr); err !=nil{
		global.Logger.Fatal("failed to start http server", zap.Error(err))
	}
}
