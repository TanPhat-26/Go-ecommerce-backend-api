package main

import (
	"fmt"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/initialize"
)

func main() {
	initialize.LoadConfig()

	fmt.Println("App Name:", global.Config.App.Name)
	fmt.Println("App Env:", global.Config.App.Env)
	fmt.Println("App Port:", global.Config.App.Port)
	fmt.Println("DB DSN:", global.Config.DB.DSN())
	fmt.Println("Redis Addr:", global.Config.Redis.Addr())
	fmt.Println("JWT Access TTL:", global.Config.JWT.AccessTTLMinutes)
	fmt.Println("Logger Level:", global.Config.Logger.Level)
}
