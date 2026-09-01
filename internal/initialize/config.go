package initialize

import (
	"log"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/setting"
)

func LoadConfig() {
	config, err := setting.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	global.Config = config
}
