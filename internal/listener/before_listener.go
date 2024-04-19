package listener

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Before() gin.HandlerFunc {
	return func(c *gin.Context) {
		trace := uuid.New().String()
		// c.Set("traceId", trace)
		c.Header("X-Request-Id", trace)

		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "origin")
		// c.Header("Content-Security-Policy", "default-src 'self';")

		c.Next()
	}
}
