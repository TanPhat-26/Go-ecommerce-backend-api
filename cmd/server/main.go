package main

import (
	"fmt"
	"log"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/setting"
)

func main() {
	config, err := setting.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("App Name:", config.App.Name)
	fmt.Println("App Env:", config.App.Env)
	fmt.Println("App Port:", config.App.Port)
	fmt.Println("DB DSN:", config.DB.DSN())
	fmt.Println("Redis Addr:", config.Redis.Addr())
	fmt.Println("JWT Access TTL:", config.JWT.AccessTTLMinutes)
	fmt.Println("Logger Level:", config.Logger.Level)
}
