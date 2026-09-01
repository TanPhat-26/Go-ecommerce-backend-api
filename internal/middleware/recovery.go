package middleware

import (
	"net/http"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc{
	return gin.CustomRecovery(func(c *gin.Context, recovered any){
		requestID, _ := c.Get(RequestIDKey)

		global.Logger.Error(
			"panic recoverd",
			zap.Any("error", recovered),
			zap.Any("request_id", requestID),
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "internal server error",
			"data": nil,
		})
	})
}