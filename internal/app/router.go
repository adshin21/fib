package app

import (
	"github.com/adshin21/fib/internal/middleware"
	"github.com/adshin21/fib/pkg/logger"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.Use(middleware.CustomGinLogger())
	router.Use(gin.Recovery())
	router.Use(middleware.UseCors())
	router.Use(middleware.RequestIDMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		l := logger.Get()
		l.Info().Msg("Hello")
		l.Debug().Msg("how are you?")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
