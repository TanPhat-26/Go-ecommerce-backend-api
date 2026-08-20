package initialize

import (
	"log"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	zapLogger, err := logger.NewLogger(global.Config.Logger)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	global.Logger = zapLogger
}
