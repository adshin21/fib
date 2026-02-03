package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Capture start time
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// 2. Process request
		c.Next()

		// 3. Calculate dimensions after request is processed
		timestamp := time.Now()
		latency := timestamp.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		reqID, ok := c.Get("RequestID")
		if !ok {
			reqID = "no-id"
		}

		fmt.Printf("[GIN] %s | %3d | %13v | %15s | %s | %-7s %#v\n",
			timestamp.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			reqID,
			method,
			path,
		)
	}
}
