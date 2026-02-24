package app

import (
	"github.com/adshin21/fib/config"
	"github.com/adshin21/fib/internal/middleware"
	"github.com/adshin21/fib/pkg/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine, cfg *config.AppConfig) {
	router.Use(middleware.CustomGinLogger())
	router.Use(gin.Recovery())
	router.Use(middleware.UseCors(cfg))
	router.Use(middleware.RequestIDMiddleware())

	l := logger.Get()
	l.Info().Str("env", cfg.Server.Env).Msg("Configured environment")

	// Debug routes with pprof (add auth middleware later)
	// Example: debugGroup.Use(authMiddleware)
	pprof.RouteRegister(router)

	router.GET("/ping", func(c *gin.Context) {
		l := logger.Get()
		l.Info().Msg("Hello")
		l.Debug().Msg("how are you?")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
}
