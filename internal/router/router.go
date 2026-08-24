package router

import (
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/controller"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.RequestID())
	r.Use(middleware.Recovery())
	r.Use(middleware.AcessLog())

	healthController := controller.NewHealthController()
	r.GET("/health", healthController.Check)

	return r
}
