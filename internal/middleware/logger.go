package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/adshin21/fib/pkg/logger"
)

func CustomGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		l := logger.Get()

		defer func() {
			latency := time.Since(start)
			statusCode := c.Writer.Status()

			if raw != "" {
				path = path + "?" + raw
			}

			reqID := c.GetString("RequestID")
			if len(reqID) == 0 {
				reqID = "no-id"
			}

			// 2. Decide on the log level based on status code
			logEvent := l.Info()
			if statusCode >= 500 {
				logEvent = l.Error()
			} else if statusCode >= 400 {
				logEvent = l.Warn()
			}

			logEvent.
				Int("code", statusCode).
				Dur("latency", latency).
				Str("client_ip", clientIP).
				Str("user_agent", userAgent).
				Str("rid", reqID).
				Str("method", method).
				Str("path", path).
				Send()
		}()

		c.Next()
	}
}
