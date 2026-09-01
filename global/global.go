package global

import (
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Config *setting.Config
var Logger *zap.Logger
var DB *gorm.DB
var Redis *redis.Client
