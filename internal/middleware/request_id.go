// Package middleware consists gin middleware
package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/adshin21/fib/internal/util"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestIDKey := "X-Request-ID"
		reqID, err := util.GetUUIDString()
		if err != nil {
			reqID = util.GenerateFastString(12)
		}
		c.Set("RequestID", reqID)
		c.Writer.Header().Set(requestIDKey, reqID)
		c.Next()
	}
}
