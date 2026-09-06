package initialize

import (
	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/controller"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/repo"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/router"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	userRepository := repo.NewUserRepository(global.DB)
	registrationRepository := repo.NewRegistrationRepository(global.DB)
	authService := service.NewAuthService(userRepository, registrationRepository)
	authController := controller.NewAuthController(authService)

	return router.NewRouter(authController)
}
