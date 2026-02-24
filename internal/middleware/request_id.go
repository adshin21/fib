// Package middleware consists gin middleware
package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/adshin21/fib/internal/util"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestIDKey := "X-Request-ID"
		var reqID string
		if s := c.Request.Header.Get(requestIDKey); len(s) > 0 {
			reqID = s
		} else {
			var err error
			reqID, err = util.GetUUIDString()
			if err != nil {
				reqID = util.GenerateFastString(12)
			}
		}
		c.Set("RequestID", reqID)
		c.Writer.Header().Set(requestIDKey, reqID)
		c.Next()
	}
}
