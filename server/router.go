package server

import (
	"github.com/adshin21/fib/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.CustomGinLogger())
	router.Use(gin.Recovery())
	router.Use(useCors())
	router.Use(middleware.RequestIDMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router
}
