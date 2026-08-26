package initialize

import (
	"fmt"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/database"
)

func InitDatabase() error {
	db, err := database.NewPostgresConnection(global.Config.DB.DSN())
	if err != nil {
		return fmt.Errorf("initialize postgres connection: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get postgres sql connection: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("ping postgres: %w", err)
	}

	global.DB = db
	return nil
}
