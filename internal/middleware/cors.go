package middleware

import (
	"time"

	"github.com/adshin21/fib/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCors(cfg *config.AppConfig) gin.HandlerFunc {
	allowOrigins := cfg.CORS.AllowOrigins
	return cors.New(
		cors.Config{
			AllowOrigins:     allowOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
}
