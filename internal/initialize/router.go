package initialize

import (
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	return router.NewRouter()
}
