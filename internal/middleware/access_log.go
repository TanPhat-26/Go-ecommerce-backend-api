package middleware

import (
	"time"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AcessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		requestID, _ := c.Get(RequestIDKey)

		fields := []zap.Field{
			zap.String("request_id", requestIDString(requestID)),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
		}

		if global.Logger == nil {
			return
		}
		switch {
		case status >= 500:
			global.Logger.Error("http request completed", fields...)
		case status >= 400:
			global.Logger.Warn("http request completed", fields...)
		default:
			global.Logger.Info("http request completed", fields...)
		}
	}
}

func requestIDString(value any) string {
	requestID, ok := value.(string)
	if !ok {
		return ""
	}
	return requestID
}
