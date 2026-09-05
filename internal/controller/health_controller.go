package controller

import (
	"net/http"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "service is healthy",
		"data": gin.H{
			"app": global.Config.App.Name,
			"env": global.Config.App.Env,
		},
	})
}
